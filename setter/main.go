package main

import (
	"time"

	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type StorageData struct {
	Data1 string `json:"data1"`
	Data2 int    `json:"data2"`
}

func main() {
	r := gin.Default()

	// redis server
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	// cache
	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	r.GET("/ping", func(ctx *gin.Context) {

		var obj string

		err := mycache.Get(ctx, "test", &obj)

		if err != nil {
			log.Printf("Error getting the cache: %v\n", err)
		}

		// rien dans le cache
		if obj == "" {
			// init de la data, puis transformation en json
			json_obj, err := json.Marshal(&StorageData{
				Data1: "value_data1",
				Data2: 3,
			})
			if err != nil {
				log.Printf("Error marshalling when initalizing object: %v\n", err)
			}

			// transformation en string pour stocakge
			obj = string(json_obj)
			// on met dans le cache (dans cache local + redis)
			err = mycache.Set(&cache.Item{
				Ctx:   ctx,
				Key:   "test",
				TTL:   time.Hour,
				Value: obj,
			})
			if err != nil {
				log.Printf("Error setting the cache: %v\n", err)
			}
		}

		ctx.JSON(200, gin.H{
			"message": obj,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package main

import (
	"net/http"
	"time"

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

	r.GET("/ping2", func(ctx *gin.Context) {

		var obj string

		err := mycache.Get(ctx, "test", &obj)

		if err != nil {
			log.Printf("Error getting the cache: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			log.Printf("Obj: %v\n", obj)

			ctx.JSON(200, gin.H{
				"message": obj,
			})
		}

	})
	r.Run(":3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

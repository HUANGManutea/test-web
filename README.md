# Test-web

POC de caching en Go en utilisant go-redis/cache, cette lib permet d'avoir un cache local et un cache redis (appel cache local puis cache redis si le cache local ne contient pas la data sohuhaitée).

# Prérequis

- installer golang
- installer podman

installer minio
```
podman run -p 9000:9000 -p 9001:9001  quay.io/minio/minio server /data --console-address ":9001"
```

installer redis
```
podman run -d --name redis_server -p 6379:6379 docker.io/redis
```

# Lancement

Terminal 1
```
cd setter
go run main.go
```

Terminal 2
```
cd getter
go run main.go
```

# Utilité du caching

setter
```
[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2021/11/11 - 21:39:00 | 200 |      2.7681ms |             ::1 | GET      "/ping"
[GIN] 2021/11/11 - 21:39:04 | 200 |        37.3µs |             ::1 | GET      "/ping     <--- récupération dans le cache/redis
```

getter
```
[GIN-debug] GET    /ping2                    --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :3001
[GIN] 2021/11/11 - 21:26:54 | 200 |      2.5916ms |             ::1 | GET      "/ping2"
[GIN] 2021/11/11 - 21:26:55 | 200 |          26µs |             ::1 | GET      "/ping2" <--- récupération sur redis
```
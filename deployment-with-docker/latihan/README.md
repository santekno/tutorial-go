# Dengan Dockerfile
```bash
$ docker build --build-arg APP_PORT=1234 --build-arg MYSQL_DBURL="root:root@tcp(127.0.0.1:3306)/course2?charset=utf8mb4&parseTime=True&loc=Local" --tag ed_app:latest -f Jatmika-Teja-Dockerfile .
$ docker run --network=host ed_app
```

# Dengan compose script
```bash
$ docker-compose -f Jatmika-Teja-docker-compose.yaml up --build
$ docker-compose -f Jatmika-Teja-docker-compose.yaml down -v
```

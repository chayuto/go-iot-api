# go-iot-api

basic API server with GO

github.com/chayuto/go-iot-api

```
go mod init go-iot-api
go run .\main.go
go get -u github.com/gin-gonic/gin

go get -u ./...
go mod tidy
```

## docker stuff

```
docker build --tag go-iot-api .
docker run --publish 8080:8080 go-iot-api

docker run -d --publish 8080:8080 go-iot-api 
docker ps
docker kill $(docker ps -q)
docker rm $(docker ps -a -q)

docker rmi $(docker images -q)
```

REF:
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
- https://github.com/rahmanfadhil/gin-bookstore

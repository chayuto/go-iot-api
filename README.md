go mod init go-iot-api
go run .\main.go
go get -u github.com/gin-gonic/gin

go mod tidy

docker build --tag go-iot-api .
docker run --publish 8080:8080 go-iot-api
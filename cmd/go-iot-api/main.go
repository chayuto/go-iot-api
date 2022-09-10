package main

import (
	"fmt"
	"net/http"

	"github.com/chayuto/go-iot-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("APP INIT")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "go-iot-api",
		})
	})

	r.GET("/measurements", controllers.FindMeasurements)
	r.GET("/measurements/:id", controllers.FindMeasurement)
	r.POST("/measurements", controllers.CreateMeasurement)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

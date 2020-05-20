package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pankaj070888/test-microservices/src/jobs/models"
)

var Job []models.Job

type Controller struct{}

func main() {
	fmt.Println("JOBS Services!!")

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	v1 := router.Group("/v1")
	{

		v1.GET("/jobs")
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":8081")
}

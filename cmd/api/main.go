package main

import (
	"middleearth/eateries/api"

	"github.com/gin-gonic/gin"
)

func ginRun() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/restaurants", api.GetRestaurants)

	r.Run() // listen and serve on 0.0.0.0:8080
}
func main() {
	ginRun()
}

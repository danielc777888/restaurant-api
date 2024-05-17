package main

import (
	"middleearth/eateries/api"
	"middleearth/eateries/data"

	"github.com/gin-gonic/gin"

	docs "middleearth/eateries/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ginRun() {
	r := gin.Default()
	db := data.Connection()
	restaurantAPI := api.NewRestaurantAPI(db)
	dishAPI := api.NewDishAPI(db)
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/restaurants", restaurantAPI.GetRestaurants)

		v1.POST("/dishes", dishAPI.CreateDish)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
func main() {
	ginRun()
}

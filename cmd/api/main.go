package main

import (
	"middleearth/eateries/api"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ginRun() {
	r := gin.Default()
	//docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/restaurants", api.GetRestaurants)
	}
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run() // listen and serve on 0.0.0.0:8080
}
func main() {
	ginRun()
}

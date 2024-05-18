package main

import (
	"context"
	"log"
	"middleearth/eateries/api"
	"middleearth/eateries/cache"
	"middleearth/eateries/data"
	"os"

	docs "middleearth/eateries/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ginRun() {
	r := gin.Default()
	db := data.Connection()

	redisAddress := os.Getenv("REDIS_ADDRESS")
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// rdb.FlushAll(ctx).Result()
	dishCache := cache.NewDishCache(rdb, &ctx)

	restaurantAPI := api.NewRestaurantAPI(db)
	dishAPI := api.NewDishAPI(db, dishCache)
	ratingAPI := api.NewRatingAPI(db)
	userAPI := api.NewUserAPI(db)
	authAPI := api.NewAuthAPI(db)
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/restaurants", restaurantAPI.GetRestaurants)

		v1.GET("/dishes/:id", dishAPI.GetDish)
		v1.GET("/dishes", dishAPI.ListDish)
		v1.POST("/dishes", dishAPI.CreateDish)
		v1.PATCH("/dishes", dishAPI.UpdateDish)
		v1.DELETE("/dishes/:id", authAPI.Authenticate, dishAPI.DeleteDish)

		v1.POST("/ratings", ratingAPI.CreateRating)

		v1.POST("/users/register", userAPI.RegisterUser)
		v1.POST("/users/login", userAPI.LoginUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
func main() {
	loadEnv()
	ginRun()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

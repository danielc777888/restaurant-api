package main

import (
	"context"
	"middleearth/eateries/api"
	"middleearth/eateries/cache"
	"middleearth/eateries/data"
	"middleearth/eateries/env"
	"middleearth/eateries/service"

	docs "middleearth/eateries/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	gpmiddleware "github.com/carousell/gin-prometheus-middleware"
)

func ginRun() {
	r := gin.Default()

	// init prometheus instrumenting middleware
	p := gpmiddleware.NewPrometheus("gin")
	p.Use(r)

	// init database
	db := data.Connection()

	// init redis
	redisAddress := env.RedisAddress()
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// rdb.FlushAll(ctx).Result()

	// init restaurant
	restaurantData := data.NewRestaurantData(db)
	restaurantService := service.NewRestaurantService(restaurantData)
	restaurantAPI := api.NewRestaurantAPI(restaurantService)

	// init dish
	dishCache := cache.NewDishCache(rdb, &ctx)
	dishAPI := api.NewDishAPI(db, dishCache)

	// init rating
	ratingData := data.NewRatingData(db)
	ratingService := service.NewRatingService(ratingData)
	ratingAPI := api.NewRatingAPI(ratingService)

	// init user
	userAPI := api.NewUserAPI(db)
	authAPI := api.NewAuthAPI(db)

	// TODO: Add auth, perhaps org middle ware, try different groupongs
	// init routes
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/restaurants", restaurantAPI.ListRestaurants)

		v1.GET("/dishes/:id", dishAPI.GetDish)
		v1.GET("/dishes", dishAPI.ListDish)
		v1.POST("/dishes", dishAPI.CreateDish)
		v1.PATCH("/dishes", dishAPI.UpdateDish)
		v1.DELETE("/dishes/:id", authAPI.Authenticate([]string{"write_dish"}), dishAPI.DeleteDish)

		v1.POST("/ratings", authAPI.Authenticate([]string{}), ratingAPI.CreateRating)

		v1.POST("/users/register", userAPI.RegisterUser)
		v1.POST("/users/login", userAPI.LoginUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	loadEnv()
	ginRun()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

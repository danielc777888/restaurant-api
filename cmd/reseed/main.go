package main

import (
	"fmt"

	"middleearth/eateries/data"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Reseeding database")
	dsn := "host=localhost user=dancingponysvc password=password dbname=dancingpony port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.Migrator().DropTable(&data.Dish{})
	db.Migrator().DropTable(&data.Restaurant{})

	db.Migrator().CreateTable(&data.Restaurant{})
	db.Migrator().CreateTable(&data.Dish{})

	db.Create(&data.Restaurant{
		Name: "The Orc Shack",
		Dishes: []data.Dish{
			{Name: "Caviar", Description: "Fish eggs", Price: 11},
			{Name: "Burger and Fries", Description: "Big portions", Price: 22},
		}})

	db.Create(&data.Restaurant{
		Name: "Dwarf Diner",
		Dishes: []data.Dish{
			{Name: "Cheese platter", Description: "Variety of cheeses from middle earth", Price: 50},
			{Name: "Sourdough", Description: "Handmade sourdough bread", Price: 40},
		}})

}

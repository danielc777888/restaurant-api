package main

import (
	"fmt"

	"middleearth/eateries/data"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Reseeding database...")
	dsn := "host=localhost user=dancingponysvc password=password dbname=dancingpony port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.Migrator().DropTable(&data.Rating{})
	db.Migrator().DropTable(&data.Dish{})
	db.Migrator().DropTable(&data.Restaurant{})
	db.Migrator().DropTable(&data.User{})

	db.Migrator().CreateTable(&data.Restaurant{})
	db.Migrator().CreateTable(&data.Dish{})
	db.Migrator().CreateTable(&data.Rating{})
	db.Migrator().CreateTable(&data.User{})

	fmt.Println("Seeding Restaurant data...")
	db.Create(&data.Restaurant{
		Name: "The Orc Shack",
		Dishes: []data.Dish{
			{Name: "Caviar", Description: "Fish eggs, tasty morsels", Price: 11},
			{Name: "Burger and Fries", Description: "Big portions of oily food", Price: 22,
				Ratings: []data.Rating{
					{Description: "This is amazing stuff"},
					{Description: "This is gross"},
				},
			},
		}})

	db.Create(&data.Restaurant{
		Name: "Dwarf Diner",
		Dishes: []data.Dish{
			{Name: "Cheese platter", Description: "Variety of cheeses from middle earth", Price: 50,
				Ratings: []data.Rating{
					{Description: "Rotten and gross"},
					{Description: "Amazingly tasty"},
				}},
			{Name: "Sourdough", Description: "Elvin sourdough bread", Price: 40},
		}})

	db.Create(&data.User{
		Name:          "Keely",
		EmailAddress:  "keely@erebor.com",
		Password:      "password",
		Locked:        false,
		LoginAttempts: 0,
	})

}

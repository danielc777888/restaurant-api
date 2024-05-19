package main

import (
	"fmt"
	"os"

	"middleearth/eateries/data"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	fmt.Println("Reseeding database...")
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.Migrator().DropTable(&data.Rating{})
	db.Migrator().DropTable(&data.Dish{})
	db.Migrator().DropTable(&data.Restaurant{})
	db.Migrator().DropTable(&data.UserPermission{})
	db.Migrator().DropTable(&data.Permission{})
	db.Migrator().DropTable(&data.User{})

	db.Migrator().CreateTable(&data.Restaurant{})
	db.Migrator().CreateTable(&data.Dish{})
	db.Migrator().CreateTable(&data.Rating{})
	db.Migrator().CreateTable(&data.User{})
	db.Migrator().CreateTable(&data.Permission{})
	db.Migrator().CreateTable(&data.UserPermission{})

	fmt.Println("Seeding Restaurant data...")

	restaurantID1, _ := uuid.Parse("e814691f-b53e-45c4-8253-e2f2a7f5ff35")
	caviarID, _ := uuid.Parse("44185708-b33e-4781-918b-baba7878348f")
	db.Create(&data.Restaurant{
		ID:   restaurantID1,
		Name: "The Orc Shack",
		Dishes: []data.Dish{
			{ID: caviarID, Name: "Caviar", Description: "Fish eggs, tasty morsels", Price: 11, RestaurantID: restaurantID1},
			{ID: uuid.New(), Name: "Burger and Fries", Description: "Big portions of oily food", Price: 22, RestaurantID: restaurantID1,
				Ratings: []data.Rating{
					{ID: uuid.New(), Description: "This is amazing stuff", RestaurantID: restaurantID1},
					{ID: uuid.New(), Description: "This is gross", RestaurantID: restaurantID1},
				},
			},
		}})

	restaurantID2, _ := uuid.Parse("522c03dc-45f6-4e74-ab28-1e882ccf74a1")
	cheesePlatterID, _ := uuid.Parse("c6625af7-4501-4592-99b4-7a5850275e44")
	db.Create(&data.Restaurant{
		ID:   restaurantID2,
		Name: "Dwarf Diner",
		Dishes: []data.Dish{
			{ID: cheesePlatterID, Name: "Cheese platter", Description: "Variety of cheeses from middle earth", Price: 50, RestaurantID: restaurantID2,
				Ratings: []data.Rating{
					{ID: uuid.New(), Description: "Rotten and gross", RestaurantID: restaurantID2},
					{ID: uuid.New(), Description: "Amazingly tasty", RestaurantID: restaurantID2},
				}},
			{ID: uuid.New(), Name: "Sourdough", Description: "Elvin sourdough bread", Price: 40, RestaurantID: restaurantID2},
		}})

	fmt.Println("Seeding User data...")

	// permissions
	adminPermission := data.Permission{
		ID:  uuid.New(),
		Key: "admin",
	}
	db.Create(&adminPermission)

	writeDishPermission := data.Permission{
		ID:  uuid.New(),
		Key: "write_dish",
	}
	db.Create(&writeDishPermission)

	db.Create(&data.Permission{
		ID:  uuid.New(),
		Key: "write_restaurant",
	})

	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	keelyID := uuid.New()
	db.Create(&data.User{
		ID:            keelyID,
		Name:          "Keely",
		EmailAddress:  "keely@erabor.com",
		Password:      string(hashedPassword1),
		Locked:        false,
		LoginAttempts: 0,
		UserPermissions: []data.UserPermission{
			{ID: uuid.New(), UserID: keelyID, PermissionID: adminPermission.ID},
		},
	})
	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("superpassword"), 10)
	gimliID := uuid.New()
	db.Create(&data.User{
		ID:            gimliID,
		Name:          "Gimli",
		EmailAddress:  "gimli@erabor.com",
		Password:      string(hashedPassword2),
		Locked:        false,
		LoginAttempts: 0,
		UserPermissions: []data.UserPermission{
			{ID: uuid.New(), UserID: gimliID, PermissionID: writeDishPermission.ID, RestaurantID: &restaurantID1},
		},
	})

}

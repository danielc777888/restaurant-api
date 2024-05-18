package api

import (
	"fmt"
	"middleearth/eateries/data"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

// requests
type RegisterUser struct {
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type LoginUser struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

// response
type LoggedInUser struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	Token        string `json:"token"`
}

type UserAPI struct {
	Db *gorm.DB
}

func NewUserAPI(Db *gorm.DB) *UserAPI {
	return &UserAPI{Db: Db}
}

// @BasePath /api/v1

func (userAPI *UserAPI) RegisterUser(c *gin.Context) {
	var user RegisterUser
	if err := c.BindJSON(&user); err != nil {
		return
	}

	// validate
	// email address, password
	// unique: email address

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	dbUser := mapRegisterUserToDB(user)
	dbUser.Password = string(hashedPassword)

	result := userAPI.Db.Create(&dbUser)
	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
	c.IndentedJSON(http.StatusOK, dbUser.ID)
}

func (userAPI *UserAPI) LoginUser(c *gin.Context) {
	var user LoginUser
	if err := c.BindJSON(&user); err != nil {
		return
	}
	var dbUser data.User
	// hash password

	// find in db
	userAPI.Db.Where("email_address = ?", user.EmailAddress).First(&dbUser)

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	jwtSecret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// dbUser.Token = &token
	// now := time.Now()
	// dbUser.TokenCreatedAt = &now
	// // save token in db
	// result := userAPI.Db.Save(&dbUser)

	//fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
	c.IndentedJSON(http.StatusOK, mapUserToJSON(dbUser, signedToken))
}

// func (dishApi *DishAPI) UpdateDish(c *gin.Context) {
// 	var dish UpdateDish
// 	if err := c.BindJSON(&dish); err != nil {
// 		fmt.Printf("Bind error %s", err)
// 		return
// 	}
// 	dbDish := mapUpdateDishToDB(dish)
// 	result := dishApi.Db.Save(&dbDish)
// 	fmt.Printf("DB result error %s, rows %d", result.Error, result.RowsAffected)
// 	c.IndentedJSON(http.StatusOK, dbDish)
// }

// func (dishApi *DishAPI) DeleteDish(c *gin.Context) {
// 	id := c.Param("id")
// 	var dish data.Dish
// 	result := dishApi.Db.First(&dish, id)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
// 		return
// 	}
// 	dishApi.Db.Delete(&dish)
// 	c.IndentedJSON(http.StatusOK, id)
// }

// func (dishApi *DishAPI) GetDish(c *gin.Context) {
// 	id := c.Param("id")
// 	var dish data.Dish
// 	result := dishApi.Db.First(&dish, id)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Dish not found"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, dish)
// }

// func (dishApi *DishAPI) ListDish(c *gin.Context) {
// 	var dishes []data.Dish
// 	dishApi.Db.Find(&dishes)
// 	c.IndentedJSON(http.StatusOK, mapDishesToJSON(dishes))
// }

func mapRegisterUserToDB(user RegisterUser) data.User {
	return data.User{
		Name:         user.Name,
		EmailAddress: user.EmailAddress,
		Password:     user.Password,
	}
}

func mapUserToJSON(user data.User, token string) LoggedInUser {
	return LoggedInUser{
		ID:           user.ID,
		Name:         user.Name,
		EmailAddress: user.EmailAddress,
		Token:        token,
	}
}

// func mapUpdateDishToDB(dish UpdateDish) data.Dish {
// 	return data.Dish{
// 		ID:           dish.ID,
// 		Name:         dish.Name,
// 		Description:  dish.Description,
// 		Price:        dish.Price,
// 		RestaurantID: dish.RestaurantID,
// 	}
// }

// func mapDishesToJSON(dishes []data.Dish) []Dish {
// 	var result []Dish
// 	for _, dish := range dishes {
// 		result = append(result, mapDishToJSON(dish))
// 	}
// 	return result
// }

// func mapDishToJSON(dish data.Dish) Dish {
// 	return Dish{
// 		Name:         dish.Name,
// 		Description:  dish.Description,
// 		Price:        dish.Price,
// 		RestaurantID: dish.RestaurantID,
// 	}
// }

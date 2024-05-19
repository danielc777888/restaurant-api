package api

import (
	"errors"
	"fmt"
	"middleearth/eateries/data"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

// requests
type RegisterUser struct {
	Name         string `json:"name" binding:"required"`
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}

type LoginUser struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
}

// response
type LoggedInUser struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Token        string `json:"token" binding:"required"`
}

type UserAPI struct {
	Db *gorm.DB
}

func NewUserAPI(Db *gorm.DB) *UserAPI {
	return &UserAPI{Db: Db}
}

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
	result := userAPI.Db.Where("email_address = ?", user.EmailAddress).First(&dbUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if dbUser.Locked {
		fmt.Println("User account locked: ", dbUser.ID)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User account locked, please contact support",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
		loginAttempt(dbUser, userAPI)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.ID.String(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	jwtSecret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		loginAttempt(dbUser, userAPI)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, mapUserToJSON(dbUser, signedToken))
}

func loginAttempt(dbUser data.User, userAPI *UserAPI) {
	dbUser.LoginAttempts = dbUser.LoginAttempts + 1
	if dbUser.LoginAttempts >= 3 {
		dbUser.Locked = true
	}
	userAPI.Db.Save(&dbUser)
}

func mapRegisterUserToDB(user RegisterUser) data.User {
	return data.User{
		ID:           uuid.New(),
		Name:         user.Name,
		EmailAddress: user.EmailAddress,
		Password:     user.Password,
	}
}

func mapUserToJSON(user data.User, token string) LoggedInUser {
	return LoggedInUser{
		ID:           user.ID.String(),
		Name:         user.Name,
		EmailAddress: user.EmailAddress,
		Token:        token,
	}
}

package service

import (
	"errors"
	"fmt"
	"middleearth/eateries/data"
	"middleearth/eateries/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// requests
type RegisterUserAction struct {
	Name         string
	EmailAddress string
	Password     string
}

type LoginUserAction struct {
	EmailAddress string
	Password     string
}

type UserResult struct {
	ID           uuid.UUID
	Name         string
	EmailAddress string
}

type LoggedInUserResult struct {
	ID           uuid.UUID
	Name         string
	EmailAddress string
	Token        string
}

type UserService struct {
	Data *data.UserData
}

func NewUserService(Data *data.UserData) *UserService {
	return &UserService{Data: Data}
}

// Registers a new user.
func (service *UserService) RegisterUser(action RegisterUserAction) (*UserResult, error) {

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(action.Password), 10)
	if err != nil {
		return nil, err
	}

	// map to user data
	user := data.User{
		ID:           uuid.New(),
		Name:         action.Name,
		EmailAddress: action.EmailAddress,
		Password:     action.Password,
	}
	user.Password = string(hashedPassword)

	createdUser, err := service.Data.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// map to user result
	result := UserResult{
		ID:           createdUser.ID,
		Name:         createdUser.Name,
		EmailAddress: createdUser.EmailAddress,
	}
	return &result, nil
}

func (service *UserService) LoginUser(action LoginUserAction) (*LoggedInUserResult, error) {

	var user data.User

	// find in db
	retrievedUser, err := service.Data.GetUserByEmailAddress(action.EmailAddress, user)
	if err != nil {
		return nil, err
	}

	if retrievedUser.Locked {
		fmt.Println("User account locked: ", retrievedUser.ID)
		return nil, errors.New("user account locked, please contact support")
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(action.Password))

	if compareErr != nil {
		fmt.Println("Error comparing password, userID: ", retrievedUser.ID)
		loginAttempt(retrievedUser)
		service.Data.UpdateUser(*retrievedUser)
		return nil, errors.New("invalid email address or password")
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": retrievedUser.ID.String(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	jwtSecret := env.JWTSecret()
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Println("Error comparing password, userID: ", retrievedUser.ID)
		loginAttempt(retrievedUser)
		fmt.Println("ERR2: ", retrievedUser)
		service.Data.UpdateUser(*retrievedUser)
		return nil, errors.New("invalid email address or password")
	}

	result := LoggedInUserResult{
		ID:           retrievedUser.ID,
		Name:         retrievedUser.Name,
		EmailAddress: retrievedUser.EmailAddress,
		Token:        signedToken,
	}
	return &result, nil
}

// A login attempt from user. If above a certain threshold, user account will be locked.
func loginAttempt(user *data.User) {
	user.LoginAttempts = user.LoginAttempts + 1
	if user.LoginAttempts >= 3 {
		user.Locked = true
	}
}

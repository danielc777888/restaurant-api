package api

import (
	"errors"
	"fmt"
	"middleearth/eateries/data"
	"middleearth/eateries/env"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthAPI struct {
	Db *gorm.DB
}

func NewAuthAPI(Db *gorm.DB) *AuthAPI {
	return &AuthAPI{Db: Db}
}

// Middleware for authenticating user and checking required permissions for an api
func (authAPI *AuthAPI) Authenticate(permissions []string) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Get the cookie off the request
		signedToken, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Validate signed token
		token, _ := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			jwtSecret := env.JWTSecret()
			return []byte(jwtSecret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the expiry date
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Find the user with token Subject
			var user data.User
			subString, _ := claims["sub"].(string)
			userId, _ := uuid.Parse(subString)
			authAPI.Db.First(&user, userId)
			fmt.Println("Authenticated user:", user.ID)
			if user.ID == uuid.Nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			if user.Locked {
				fmt.Println("User account locked:", user.ID)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			if len(permissions) > 0 {
				// get permissions from user
				var userPermissions []data.UserPermission
				authAPI.Db.Joins("Permission", authAPI.Db.Where("user_id = ?", user.ID)).Find(&userPermissions)

				if !isSuperAdmin(userPermissions) {
					restaurantID, _ := GetRestaurantHeader(c)
					if !hasAllPermissions(restaurantID, userPermissions, permissions) {
						c.AbortWithStatus(http.StatusUnauthorized)
						return
					}
				}
			}

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// Checks to see if users has all the required permissions, using her userPermissions list.
// TODO: Use a map instead of nested loops
func hasAllPermissions(restaurantID uuid.UUID, userPermissions []data.UserPermission, permissions []string) bool {
	for _, permission := range permissions {
		if !hasPermission(restaurantID, userPermissions, permission) {
			return false
		}
	}
	return true
}

// Checks to see whether userPermissions has a specific required permission.
func hasPermission(restaurantID uuid.UUID, userPermissions []data.UserPermission, permission string) bool {
	for _, userPermission := range userPermissions {

		// check restaurant
		if restaurantID != uuid.Nil && (userPermission.RestaurantID == nil || restaurantID != (*userPermission.RestaurantID)) {
			continue
		}

		// check permission key
		if userPermission.Permission.Key == permission {
			return true
		}
	}
	return false
}

// Checks to see if user is a superAdmin. ie. Has a admin permission not associated with a restaurant.
func isSuperAdmin(userPermissions []data.UserPermission) bool {
	for _, userPermission := range userPermissions {
		if userPermission.Permission.Key == "admin" && userPermission.RestaurantID == nil {
			return true
		}
	}
	return false
}

// Extracts auth bearer token from Auth header
func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("invalid header")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("invalid bearer token header")
	}

	return jwtToken[1], nil
}

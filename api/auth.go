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
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			jwtSecret := env.JWTSecret()
			return []byte(jwtSecret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the expiry date
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// Find the user with token Subject
			var user data.User
			subString, _ := claims["sub"].(string)
			userId, _ := uuid.Parse(subString)
			authAPI.Db.First(&user, userId)
			fmt.Println("##### AUTH User:", user.ID)
			if user.ID == uuid.Nil {
				fmt.Println("##### ABORT User:", user.ID)
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			if user.Locked {
				fmt.Println("User account locked:", user.ID)
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			// authoroise if there are required permissions
			if len(permissions) > 0 {
				// get permissions from user
				var userPermissions []data.UserPermission
				authAPI.Db.Joins("Permission", authAPI.Db.Where("user_id = ?", user.ID)).Find(&userPermissions)

				fmt.Println("### GOT USER PERMISSIONS:", userPermissions)

				if !isSuperAdmin(userPermissions) {
					restaurantID, _ := GetRestaurantHeader(c)
					if !hasAllPermissions(restaurantID, userPermissions, permissions) {
						fmt.Println("$$$$ USER DOES NOT HAVE ALL THE FOLLOWING PERMISSIONS:")
						c.AbortWithStatus(http.StatusUnauthorized)
					}
				}
			}

			// if user does not have required permissions throw an error

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// TODO Use a map to prevent nested loops
func hasAllPermissions(restaurantID uuid.UUID, userPermissions []data.UserPermission, permissions []string) bool {
	for _, permission := range permissions {
		if !hasPermission(restaurantID, userPermissions, permission) {
			return false
		}
	}
	fmt.Println(">>>>>> USER HAS ALL PERMISSIONS: GOOD TO GO!!!!")
	return true
}

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

func isSuperAdmin(userPermissions []data.UserPermission) bool {
	for _, userPermission := range userPermissions {
		if userPermission.Permission.Key == "admin" && userPermission.RestaurantID == nil {
			fmt.Println("#### FOUND SUPER ADMIN!!!!")
			return true
		}
	}
	return false
}

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

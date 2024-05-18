package api

import (
	"errors"
	"fmt"
	"log"
	"middleearth/eateries/data"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthAPI struct {
	Db *gorm.DB
}

func NewAuthAPI(Db *gorm.DB) *AuthAPI {
	return &AuthAPI{Db: Db}
}

func (authAPI *AuthAPI) Authenticate(c *gin.Context) {
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
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		var user data.User
		authAPI.Db.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if user.Locked == true {
			log.Println("User account locked:", user.ID)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the request
		c.Set("user", user)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

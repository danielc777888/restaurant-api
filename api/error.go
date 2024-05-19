package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Standard response to an error with StatusBadRequest
func ErrorResponse(err error, c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

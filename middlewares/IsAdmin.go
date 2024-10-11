package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Value("role") != "admin"{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not autorized as admin"})
			c.Abort()
			return
		}
		c.Next()
	}
}
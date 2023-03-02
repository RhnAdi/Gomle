package middleware

import (
	"net/http"

	"github.com/RhnAdi/Gomle/server/internal/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "failed",
				"message": "requst header auth empty",
			})
			c.Abort()
			return
		}
		claim, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"status": "failed", "message": err.Error()})
			c.Abort()
			return
		}
		c.Set("claim", claim)
		c.Next()
	}
}

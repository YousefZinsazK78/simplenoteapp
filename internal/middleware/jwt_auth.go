package middleware

import (
	"net/http"
	"notegin/internal/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authrozationHeader := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authrozationHeader, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}
		tokenString := strings.TrimPrefix(authrozationHeader, "Bearer ")
		claims, err := utils.GetClaimsFromJwtToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}

		var exp_time, ok = claims["ExpiresAt"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}
		var exp, okN = exp_time.(int64)
		if !okN {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}
		if time.Now().Unix() > exp {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}
		c.Set("user_id", claims["user_id"])
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

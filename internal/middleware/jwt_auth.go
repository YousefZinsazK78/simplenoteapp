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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized! has not bearer prefix"})
			return
		}
		tokenString := strings.TrimPrefix(authrozationHeader, "Bearer ")
		claims, err := utils.GetClaimsFromJwtToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}
		var exp_time, ok = claims["exp"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}
		var exp, okN = exp_time.(float64)
		if !okN {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
			return
		}
		if time.Now().Unix() > int64(exp) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!  Token Expired!!"})
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}

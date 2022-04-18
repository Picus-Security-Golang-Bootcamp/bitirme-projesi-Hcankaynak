package middleware

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/user"
	jwt_service "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwt_service.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				if user.IsAdmin(decodedClaims.Role) {
					c.Next()
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}

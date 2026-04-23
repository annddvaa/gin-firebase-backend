package middleware

import (
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(401, gin.H{"message": "Token tidak ditemukan"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(401, gin.H{"message": "Token tidak valid"})
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(401, gin.H{"message": "Gagal membaca claims"})
            return
        }

        var userID interface{}
        if val, exists := claims["sub"]; exists {
            userID = val
        }
        
        role, _ := claims["role"].(string)
        c.Set("user_id", userID)
        c.Set("role", role)
        c.Next()
    }
}
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "role tidak ditemukan",
			})
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "role tidak valid",
			})
			return
		}

		if roleStr != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "akses hanya admin",
			})
			return
		}

		c.Next()
	}
}
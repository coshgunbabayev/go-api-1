package middleware

import (
	"go-api-1/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthenticateForPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.Redirect(302, "/account")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.Redirect(302, "/account")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["userId"].(string)

			var userModel models.UserModel

			user, _ := userModel.GetByID(userId)

			if user.IsEmpty() {
				c.Redirect(302, "/account")
				c.Abort()
				return
			}

			c.Set("user", user)
		} else {
			c.Redirect(302, "/account")
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthenticateForAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(302, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(302, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["userId"].(string)

			var userModel models.UserModel

			user, _ := userModel.GetByID(userId)

			if user.IsEmpty() {
				c.JSON(302, gin.H{
					"success": false,
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}

			c.Set("user", user)
		} else {
			c.JSON(302, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// // JWT token'ı veya cookie'den token al
		// tokenString, err := c.Cookie("access_token")
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Yetkisiz erişim"})
		// 	c.Abort()
		// 	return
		// }

		// // JWT token'ı doğrula
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	// Token doğrulama anahtarını burada ayarlayın (örneğin, bir gizli anahtar)
		// 	return []byte("gizli_anahtar"), nil
		// })
		// if err != nil || !token.Valid {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz token"})
		// 	c.Abort()
		// 	return
		// }

		// // Token geçerli ise, isteği devam ettir
		// c.Next()

		aaa := c.Param("aaa")

		if aaa == "salam" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Geçersiz token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

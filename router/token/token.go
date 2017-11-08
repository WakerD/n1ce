package token

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	tokenStr, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("123456"), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	fmt.Println("GG")

	if !tokenStr.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}

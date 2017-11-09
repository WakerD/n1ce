package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewJwtToken(userID string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 168) //一周有效期
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	out, err := token.SignedString([]byte("123456"))
	return out, err
}

func GetIDFromToken(auth string) (string, error) {
	token, _ := jwt.Parse(auth, nil)
	claims := token.Claims.(jwt.MapClaims)
	id, ok := claims["sub"]
	if !ok {
		return "", errors.New("no sub")
	}
	return id.(string), nil
}

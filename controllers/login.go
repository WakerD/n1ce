package controllers

import (
	"fmt"

	"n1ce/common"
	"n1ce/system"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")

	user, err := models.UserIns.ReadOne(bson.M{"account": account})
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	if common.EncodeMessageMd5(password) == user.Password {
		token, err := NewJwtToken(user.ID)
		if err != nil {
			c.JSON(http.StatusOK, token)
		}
	}
}

func Signup(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err := models.UserIns.Create(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func NewJwtToken(userID string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["ID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 168) //一周有效期
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	out, err := token.SignedString([]bye("dsij#43jl3#$"))
	return out, err
}

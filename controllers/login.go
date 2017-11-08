package controllers

import (
	"net/http"
	"time"

	"n1ce/cache"
	"n1ce/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func Signup(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err = models.UserIns.Create(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func Signin(c *gin.Context) {
	data := models.User{}
	err := c.Bind(&data)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	user, err := models.UserIns.ReadOne(bson.M{"account": data.Account})

	if user.VerifyPassword(data.Password) {
		token, _ := NewJwtToken(string(user.ID))
		cache.RedisCli.Do("HSET", "token:"+token, "userID", user.ID.Hex())
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "XX"})
	}
}

func NewJwtToken(userID string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["ID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 168) //一周有效期
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	out, err := token.SignedString([]byte("123456"))
	return out, err
}

/*
func ObjectIdToString(id bson.ObjectId) string {
	idStr := id.String()
	res := strings.TrimPrefix(idStr, "ObjectId(\"")
	res = strings.TrimSuffix(res, ")\"")
	return res
}*/

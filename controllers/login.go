package controllers

import (
	"net/http"

	//	"n1ce/cache"
	"n1ce/models"
	"n1ce/util/token"

	//	"github.com/dgrijalva/jwt-go"
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
		token, _ := token.NewJwtToken(user.ID.Hex())
		//		cache.RedisCli.Do("HSET", "token:"+token, "userID", user.ID.Hex())
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"msg": "wrong"})
	}
}

/*
func ObjectIdToString(id bson.ObjectId) string {
	idStr := id.String()
	res := strings.TrimPrefix(idStr, "ObjectId(\"")
	res = strings.TrimSuffix(res, ")\"")
	return res
}*/

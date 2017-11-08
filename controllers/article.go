package controllers

import (
	//	"encoding/hex"
	"fmt"
	"net/http"
	//	"strconv"

	"n1ce/cache"
	"n1ce/models"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"github.com/davecgh/go-spew/spew"
)

var (
	ArticleIns = new(Article)
)

type Article struct{}

/*
func (*Article) Get(c *gin.Context) {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}*/

func (*Article) Create(c *gin.Context) {
	fmt.Println("XX")
	data := models.Article{}
	err := c.Bind(&data)
	spew.Printf(c.GetHeader("Authorization"))
	userID, err := redis.String(cache.RedisCli.Do("HGET", "token:"+c.GetHeader("Authorization"), "userID"))
	data.UserID = bson.ObjectIdHex(userID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err = models.ArticleIns.Create(data)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "sucess"})
}

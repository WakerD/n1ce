package controllers

import (
	//	"encoding/hex"
	"net/http"
	"strconv"

	"n1ce/cache"
	"n1ce/models"

	"github.com/davecgh/go-spew/spew"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
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
	data := models.Article{}
	err := c.Bind(&data)
	userID, err := redis.String(cache.RedisCli.Do("HGET", "token:"+c.GetHeader("Authorization"), "userID"))
	data.UserID = bson.ObjectIdHex(userID)
	spew.Printf("%#+v\n", data.UserID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err = models.ArticleIns.Create(data)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "sucess"})
}

func (*Article) Update(c *gin.Context) {
	data := models.Article{}
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "request data wrong"})
	}
	/*if !data.Valid() {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "request data wrong"})
	}*/
	id := c.Param("id")

	userID, err := redis.String(cache.RedisCli.Do("HGET", "token:"+c.GetHeader("Authorization"), "userID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	selector := bson.M{
		"userid": bson.ObjectIdHex(userID),
		"_id":    bson.ObjectIdHex(id),
	}

	update := bson.M{
		"$set": bson.M{"title": data.Title,
			"content": data.Content,
			"pics":    data.Pics,
		},
	}
	spew.Printf("%#+v\n", bson.ObjectIdHex(id))
	err = models.ArticleIns.Update(selector, update)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "sucess"})
}

func (*Article) ReadMany(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	userID, err := redis.String(cache.RedisCli.Do("HGET", "token:"+c.GetHeader("Authorization"), "userID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	data, err := models.ArticleIns.ReadMany(bson.M{"userid": bson.ObjectIdHex(userID)}, offset, limit)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, data)
}

func (*Article) ReadOne(c *gin.Context) {
	id := c.Param("id")

	userID, err := redis.String(cache.RedisCli.Do("HGET", "token:"+c.GetHeader("Authorization"), "userID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	selector := bson.M{
		"userid": bson.ObjectIdHex(userID),
		"_id":    bson.ObjectIdHex(id),
	}
	data, err := models.ArticleIns.ReadOne(selector)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, data)
}

func (*Article) DeleteOne(c *gin.Context) {
	id := c.Param("id")

	userID, err := redis.String(cache.RedisCli.Do("HGET", "token:"+c.GetHeader("Authorization"), "userID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	selector := bson.M{
		"userid": bson.ObjectIdHex(userID),
		"_id":    bson.ObjectIdHex(id),
	}
	cnt, err := models.ArticleIns.Delete(selector, false)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, cnt)
}

func (*Article) DeleteMany(c *gin.Context) {
	userID, err := redis.String(cache.RedisCli.Do("HGET", "token:"+c.GetHeader("Authorization"), "userID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	selector := bson.M{
		"userid": bson.ObjectIdHex(userID),
	}
	cnt, err := models.ArticleIns.Delete(selector, true)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, cnt)
}

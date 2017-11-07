package controllers

import (
	"fmt"
	"net/http"
	//	"strconv"

	"n1ce/models"

	"github.com/gin-gonic/gin"
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
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	err = models.ArticleIns.Create(data)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "sucess"})
}

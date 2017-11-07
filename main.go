package main

import (
	"n1ce/cache"
	"n1ce/controllers"
	"n1ce/models"

	"n1ce/router/header"
	//	"n1ce/system"

	"github.com/gin-gonic/gin"
)

func init() {
	connectToDB()
	cache.GetRedis()
}

func main() {
	router := gin.Default()
	router.Use(header.NoCache)
	router.Use(header.Options)
	router.Use(header.Secure)

	router.POST("/signin", controllers.Signin)
	router.POST("/signup", controllers.Signup)

	article := router.Group("/article")
	{
		router.Use()
		/*		article.GET("", controllers.ArticleIns.Get)
				article.GET("/:id", controllers.ArticleIns.GetSelf)(/)*/
		article.POST("", controllers.ArticleIns.Create)
	}
	router.Run(":3737")
}

func connectToDB() {
	//	config := system.GetConfig()
	//	connection := fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port)
	models.SetDB("localhost", "test")
}

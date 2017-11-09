package main

import (
	"n1ce/cache"
	"n1ce/controllers"
	"n1ce/models"
	"n1ce/router/header"
	"n1ce/router/token"
	"n1ce/util/logs"
	//	"n1ce/system"

	"github.com/gin-gonic/gin"
)

func init() {
	lg.Log.Infof(" Start ...")
	connectToDB()
	err := cache.GetRedis()
	if err != nil {
		lg.Log.Fatalf("get redis failed")
	}
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
		article.Use(token.VerifyToken)
		/*		article.GET("", controllers.ArticleIns.Get)
				article.GET("/:id", controllers.ArticleIns.GetSelf)(/)*/
		article.POST("", controllers.ArticleIns.Create)
		article.GET("", controllers.ArticleIns.ReadMany)
		article.DELETE("", controllers.ArticleIns.DeleteMany)
		article.PUT("/:id", controllers.ArticleIns.Update)
		article.GET("/:id", controllers.ArticleIns.ReadOne)
		article.DELETE("/:id", controllers.ArticleIns.DeleteOne)
	}
	router.Run(":3737")
}

func connectToDB() {
	//	config := system.GetConfig()
	//	connection := fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port)
	models.SetDB("localhost", "test")
}

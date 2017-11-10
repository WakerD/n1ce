package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"n1ce/server/controllers"
	"n1ce/server/router/header"
	"n1ce/server/router/token"
)

func Load() http.Handler {
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

	return router
}

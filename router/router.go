package router

import (
	"github.com/gin-gonic/gin"
)

func Load() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.Use(header.NoCache)
	r.Use(header.Options)
	r.Use(header.Secure)
	r.Use(session.SetUser)
	r.Use(token.Refresh)

	r.GET("/logout", server.GetLogout)
	r.GET("/login", server.HandleLogin)

	auth := r.Group("/authorize")
	{
		auth.GET("", server.HandleAuth)
		auth.POST("", server.HandleAuth)
		auth.POST("/token", server.GetLoginToken())
	}
}

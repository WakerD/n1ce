package main

import (
	"n1ce/controllers"
	"n1ce/models"
	"n1ce/system"

	"github.com/gin-gonic/gin"
)

func init() {
	connectToDB()
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.Post("/signin", controllers.Signin)
	router.Post("/signup", controllers.Signup)
}

func connectToDB() {
	config := system.GetDBConfig()
	connection := fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port)
	models.SetDB(connection, "test")
}

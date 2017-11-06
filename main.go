package main

import "fmt"

func main() {
	router := gin.Default()

	router.Post("/signin", controllers.Signin)
	router.Post("/signup", controllers.Signup)
}

func connectToDB() {
	config := system.GetDBConfig()
	connection := fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port)
	models.SetDB(connection, config.Mode)
}

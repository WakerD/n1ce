package main

import "fmt"

func main() {
	router := gin.Default()

	router.Post("/signin", controllers.Signin)
	router.Post("/signup", controllers.Signup)
}

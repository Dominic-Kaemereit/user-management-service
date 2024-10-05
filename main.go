package main

import (
	"github.com/Dominic-Kaemereit/user-management-service/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// create a new router
	router := gin.Default()

	// define the routes
	router.PUT("/user", controller.PostCreateUser)
	router.GET("/user", controller.GetFindUserWithNameOrEmail)
	router.GET("/user/login", controller.GetUserLogin)
	router.GET("/users", controller.GetAListFromAllUser)

	// run the server
	err := router.Run("127.0.0.1:8080")

	if err != nil {
		panic(err)
	}
}

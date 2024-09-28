package main

import (
	"github.com/Dominic-Kaemereit/user-management-service/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// create a new router
	router := gin.Default()

	// define the routes
	router.PUT("/user", controller.CreateUser)
	router.GET("/user", controller.FindUserWithNameOrEmail)
	router.GET("/users", controller.GetAListFromAllUser)

	// run the server
	router.Run("127.0.0.1:8080")
}

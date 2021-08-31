package app

import (
	"github.com/neazossa/bookstore-user-api/controllers"
	"github.com/neazossa/bookstore-user-api/controllers/users"
)

func mapURL(){
	router.GET("/ping", controllers.Ping)

	//USERS ENDPOINT
	router.GET("/users/:id", users.GetUser)
	router.GET("/users/search", users.FindUser)
	router.POST("/users/", users.CreateUser)
}

package app

import (
	"github.com/roma11111122/bookstore_users-api/controllers/ping"
	"github.com/roma11111122/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/user/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/users/controllers/ping"
	"github.com/guiaramos/bookstore/users/controllers/users"
)

func mapUrls(r *gin.Engine) {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:user_id", users.GetUser)
	r.GET("/users/search", users.FindUser)
	r.POST("/users", users.CreateUser)
}

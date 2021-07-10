package app

import (
	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/users/controllers/ping"
	"github.com/guiaramos/bookstore/users/controllers/users"
)

func mapUrls(r *gin.Engine) {
	r.GET("/ping", ping.Ping)

	r.POST("/users", users.CreateUser)
	r.GET("/users/:user_id", users.GetUser)
	r.PUT("/users/:user_id", users.UpdateUser)
	r.PATCH("/users/:user_id", users.UpdateUser)
	r.GET("/users/search", users.FindUser)
}

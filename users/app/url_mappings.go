package app

import (
	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/users/controllers/ping"
	"github.com/guiaramos/bookstore/users/controllers/users"
)

func mapUrls(r *gin.Engine) {
	r.GET("/ping", ping.Ping)

	r.POST("/users", users.Create)
	r.GET("/users/:user_id", users.Get)
	r.PUT("/users/:user_id", users.Update)
	r.PATCH("/users/:user_id", users.Update)
	r.GET("/users/search", users.Find)
	r.DELETE("/users/:user_id", users.Delete)
}

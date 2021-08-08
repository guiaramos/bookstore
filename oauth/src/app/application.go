package app

import (
	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/oauth/src/clients/cassandra"
	"github.com/guiaramos/bookstore/oauth/src/domain/access_token"
	"github.com/guiaramos/bookstore/oauth/src/http"
	"github.com/guiaramos/bookstore/oauth/src/repository/db"
)

var (
	router = gin.Default()
)

// StartApplication function starts the application
func StartApplication() {
	_ = cassandra.GetSession()

	repo := db.NewDBRepository()
	service := access_token.NewService(repo)
	handler := http.NewHandler(service)

	router.GET("/oauth/access_token/:access_token_id", handler.GetByID)
	router.POST("/oauth/access_token", handler.Create)

	router.Run(":8080")
}

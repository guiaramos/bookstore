package app

import (
	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/users/logger"
)

func StartApplication() {
	r := gin.Default()

	mapUrls(r)

	logger.Info("about to start the application...")
	r.Run(":8080")
}

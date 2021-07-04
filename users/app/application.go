package app

import (
	"github.com/gin-gonic/gin"
)

func StartApplication() {

	r := gin.Default()
	mapUrls(r)
	r.Run(":8080")
}

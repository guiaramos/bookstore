package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/users/domain/users"
	"github.com/guiaramos/bookstore/users/services"
	"github.com/guiaramos/bookstore/users/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	r, e := services.CreateUser(user)
	if e != nil {
		c.JSON(e.Status, e)
		return
	}

	c.JSON(http.StatusCreated, r)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

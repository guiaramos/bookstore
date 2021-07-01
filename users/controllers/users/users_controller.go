package users

import (
	"net/http"
	"strconv"

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
	userId, convErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if convErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

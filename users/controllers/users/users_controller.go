package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/users/domain/users"
	"github.com/guiaramos/bookstore/users/services"
	"github.com/guiaramos/bookstore/users/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, convErr := strconv.ParseInt(userIdParam, 10, 64)
	if convErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}

	return userId, nil
}

func Create(c *gin.Context) {
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

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	user, err := services.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)

}

func Find(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid body json")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

package http

import (
	"github.com/guiaramos/bookstore/oauth/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/oauth/src/domain/access_token"
)

// Handler interface defines a Handler
type Handler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type handler struct {
	service access_token.Service
}

// NewHandler function creates a new handler
func NewHandler(service access_token.Service) Handler {
	return &handler{
		service: service,
	}
}

// GetByID method handles the get of access token by its ID
func (h *handler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

// Create method handles the creation of new access token
func (h *handler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guiaramos/bookstore/oauth/src/domain/access_token"
)

// Handler interface defines a Handler
type Handler interface {
	GetByID(*gin.Context)
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

// GetByID method gets an access token by its ID
func (h *handler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

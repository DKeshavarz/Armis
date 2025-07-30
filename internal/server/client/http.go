package client

import (
	"context"
	"log"
	"net/http"

	"github.com/DKeshavarz/armis/internal/servise"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	servise servise.ServiceInterfase
}
func RegisterRoutes(group *gin.RouterGroup, servise servise.ServiceInterfase) {
	h := &Handler{servise: servise}
	group.GET("/:key", h.getValue)
}

func (h *Handler) getValue(c *gin.Context) {
	key := c.Param("key")
	log.Println("key =", key)
	ctx := context.Background()
	value, err := h.servise.Get(ctx, key)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"key": key, "value": value})
}
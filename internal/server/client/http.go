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
	group.PUT("/:key", h.putValue)
	group.DELETE("/:key", h.deleteKey)
}

func (h *Handler) getValue(c *gin.Context) {
	key := c.Param("key")
	log.Println("key =", key)
	ctx := context.Background()
	value, err := h.servise.Get(ctx, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

func (h *Handler) putValue(c *gin.Context) {
	key := c.Param("key")
	var req struct {
		Value string `json:"value"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	ctx := context.Background()
	if err := h.servise.Put(ctx, key, req.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "key stored", "key": key, "value": req.Value})
}

func (h *Handler) deleteKey(c *gin.Context) {
	key := c.Param("key")

	ctx := context.Background()
	if err := h.servise.Delete(ctx, key); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "key deleted", "key": key})
}

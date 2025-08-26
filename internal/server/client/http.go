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

// getValue godoc
// @Summary      Get a value
// @Description  Returns the value associated with a key
// @Tags         client
// @Param        key   path      string  true  "Key to look up"
// @Produce      json
// @Success      200  {object}   ValueResponse
// @Failure      400  {object}   ErrorResponse
// @Router       /client/{key} [get]
func (h *Handler) getValue(c *gin.Context) {
	key := c.Param("key")
	log.Println("key =", key)
	ctx := context.Background()
	value, err := h.servise.Get(ctx, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ValueResponse{Key: key, Value: value})
}

// putValue godoc
// @Summary      Store or update a value
// @Description  Stores a value under the given key
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        key   path      string          true  "Key to store"
// @Param        request body   PutValueRequest  true  "Request body with value"
// @Success      200  {object}  MessageResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /client/{key} [put]
func (h *Handler) putValue(c *gin.Context) {
	key := c.Param("key")
	var req PutValueRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.Value == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid input"})
		return
	}

	ctx := context.Background()
	if err := h.servise.Put(ctx, key, req.Value); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "key stored", Key: key, Value: req.Value})
}

// deleteKey godoc
// @Summary      Delete a key
// @Description  Deletes the key and its value
// @Tags         client
// @Param        key   path      string  true  "Key to delete"
// @Produce      json
// @Success      200  {object}   MessageResponse
// @Failure      404  {object}   ErrorResponse
// @Router       /client/{key} [delete]
func (h *Handler) deleteKey(c *gin.Context) {
	key := c.Param("key")

	ctx := context.Background()
	if err := h.servise.Delete(ctx, key); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "key deleted", Key: key})
}

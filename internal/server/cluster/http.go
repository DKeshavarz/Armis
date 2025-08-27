package cluster

import (
	"net/http"

	"github.com/DKeshavarz/armis/internal/cluster"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cluster cluster.Cluster
}

func RegisterRoutes(group *gin.RouterGroup) {
	handle := Handler{
		cluster: nil,
	}
	group.GET("/ping", handle.pingReply)
}

func (h *Handler) pingReply(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status" : "I am alive"})
}
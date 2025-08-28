package cluster

import (
	"net/http"

	"github.com/DKeshavarz/armis/internal/config"
	"github.com/DKeshavarz/armis/pkg/cluster"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cluster cluster.Cluster
}

func RegisterRoutes(group *gin.RouterGroup) {
	cfg, _ := config.New()
	handle := Handler{
		cluster: cluster.New(cfg.Cluster),
	}
	group.GET("/ping", handle.pingReply)
	group.POST("/join", handle.joinReply)
}

func (h *Handler) pingReply(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"status": "I am alive",
		"info" : h.cluster.ACK(),
	})
}

func (h *Handler) joinReply(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"status": "Wellcome",
		"info" : h.cluster.JoinReply(),
	})
}

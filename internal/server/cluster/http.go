package cluster

import (
	"fmt"
	"log"
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
	c.JSON(http.StatusOK, cluster.PingResponse{
		Msg: "Ping ok",
		Info: h.cluster.ACK(),
	})
}

func (h *Handler) joinReply(c *gin.Context) {
	var req cluster.JoinRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

	fmt.Println("hello" , req)
	rep := h.cluster.JoinReply()

	log.Println(rep)
	c.JSON(http.StatusOK, cluster.JoinResponse{
		Msg:  "Join ok",
		Info: rep,
	})
}

package cluster

import (
	"net/http"

	"github.com/DKeshavarz/armis/internal/logger"
	"github.com/DKeshavarz/armis/pkg/cluster"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	cluster cluster.Cluster
	logger   logger.Logger
}

func RegisterRoutes(group *gin.RouterGroup, cluster cluster.Cluster) {
	handle := Handler{
		cluster: cluster,
		logger: logger.New("cluster-handel"),
	}
	group.GET("/ping", handle.pingReply)
	group.POST("/join", handle.joinReply)
}

func (h *Handler) pingReply(c *gin.Context) {
	resp := h.cluster.ACK()
	c.JSON(http.StatusOK, cluster.PingResponse{
		Msg: "Ping ok",
		Info: resp,
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

	h.logger.Debug("get join body", logger.Field{Key: "body", Value: req})
	h.cluster.GetUpdate(req.Self)
	rep := h.cluster.JoinReply()

	c.JSON(http.StatusOK, cluster.JoinResponse{
		Msg:  "Join ok",
		Info: rep,
	})
}

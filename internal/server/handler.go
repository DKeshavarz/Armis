package server

import (
	"github.com/DKeshavarz/armis/internal/server/client"
	"github.com/DKeshavarz/armis/internal/servise"
	"github.com/gin-gonic/gin"
)

func New(servise servise.ServiceInterfase)*gin.Engine {
	r := gin.Default()
	setup(r, servise)
	return r
}

func setup(r *gin.Engine, servise servise.ServiceInterfase) {
	client.RegisterRoutes(r.Group("/client"),servise)
}
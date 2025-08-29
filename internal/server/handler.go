package server

import (
	_ "github.com/DKeshavarz/armis/docs"
	"github.com/DKeshavarz/armis/internal/server/client"
	clusterHandle "github.com/DKeshavarz/armis/internal/server/cluster"
	"github.com/DKeshavarz/armis/internal/servise"
	clu "github.com/DKeshavarz/armis/pkg/cluster"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Armis Key-Value Store API
// @version         0.8.0
// @description     Simple key-value store service with CRUD operations.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Armis Dev Team
// @contact.url    https://github.com/DKeshavarz/armis
// @contact.email  dankeshavarz1075@example.com

// @BasePath  /

// @tag.name        client
// @tag.description Operations for managing key-value pairs as a client

func New(servise servise.ServiceInterfase ,cluster clu.Cluster) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setup(r, servise, cluster)
	return r
}

func setup(r *gin.Engine, servise servise.ServiceInterfase, cluster clu.Cluster) {
	client.RegisterRoutes(r.Group("/client"), servise)
	clusterHandle.RegisterRoutes(r.Group("/cluster"), cluster)
}

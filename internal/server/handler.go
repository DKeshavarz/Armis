package server

import (
	_ "github.com/DKeshavarz/armis/docs"
	"github.com/DKeshavarz/armis/internal/server/client"
	"github.com/DKeshavarz/armis/internal/server/cluster"
	"github.com/DKeshavarz/armis/internal/servise"
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

func New(servise servise.ServiceInterfase) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setup(r, servise)
	return r
}

func setup(r *gin.Engine, servise servise.ServiceInterfase) {
	client.RegisterRoutes(r.Group("/client"), servise)
	cluster.RegisterRoutes(r.Group("/cluster"))
}

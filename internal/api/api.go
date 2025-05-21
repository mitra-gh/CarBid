package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/api/middlewares"
	"github.com/mitra-gh/CarBid/internal/api/routers"
)
var router *gin.Engine

func InitServer(cfg *configs.Config) {
	c := gin.Context{}
	router := gin.New()


	router.Use(middlewares.DefaultStructuredLogger(cfg))
	router.Use(gin.Logger(), gin.Recovery())
	

	swagger := router.Group("/swagger")
	routers.SwaggerRouter(&c, swagger,cfg)


	v1 := router.Group("/api/v1")
	routers.V1Router(&c, v1,cfg)

	v1.Use(middlewares.DefaultStructuredLogger(cfg))

	router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}


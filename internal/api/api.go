package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
)
var router *gin.Engine

func InitServer(cfg *configs.Config) {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())


	router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}


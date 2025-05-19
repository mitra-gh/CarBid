package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())


	v1 := router.Group("/api/v1")

	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
	}

	router.Run(":8080")
}
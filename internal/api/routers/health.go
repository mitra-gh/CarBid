package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/internal/api/handlers"
)

func HealthRouter(c *gin.Context, r *gin.RouterGroup) {
	healthHandler := handlers.HealthHandler{}
	r.GET("/", healthHandler.HealthCheck)
}
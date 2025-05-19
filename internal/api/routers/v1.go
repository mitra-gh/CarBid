package routers

import "github.com/gin-gonic/gin"

func V1Router(c *gin.Context, r *gin.RouterGroup) {
	health := r.Group("/health")
	HealthRouter(c, health)

}
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
)

func V1Router(c *gin.Context, r *gin.RouterGroup,cfg *configs.Config) {
	health := r.Group("/health")
	HealthRouter(c, health)

	user := r.Group("/user")
	UserRouter(c, user,cfg)

}
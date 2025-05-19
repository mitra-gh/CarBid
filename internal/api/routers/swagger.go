package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRouter(c *gin.Context, r *gin.RouterGroup,cfg *configs.Config) {
	docs.SwaggerInfo.Title = "CarBid API"
	docs.SwaggerInfo.Description = "CarBid API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + cfg.Server.Port
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	
	r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
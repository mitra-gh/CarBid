package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/api/handlers"
)

func UserRouter(c *gin.Context,router *gin.RouterGroup,cfg *configs.Config) {
	UserHandler := handlers.NewUserHandler(cfg)


	router.POST("/send-otp", UserHandler.SendOtp)


}
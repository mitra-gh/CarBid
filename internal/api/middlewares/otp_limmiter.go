package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
	responseHelper "github.com/mitra-gh/CarBid/internal/api/response-helper"
	ipLimiter "github.com/mitra-gh/CarBid/pkg/ip_limiter"
	serviceErrors "github.com/mitra-gh/CarBid/pkg/service_errs"
	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *configs.Config) gin.HandlerFunc {
	ipLimiter := ipLimiter.NewIPRateLimiter(rate.Limit(cfg.Otp.Limiter), 5)
	return func(c *gin.Context) {
		limiter := ipLimiter.GetLimiter(c.Request.RemoteAddr)
		if limiter != nil && !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, responseHelper.Response{
				Code:    -2,
				Message: serviceErrors.ErrTooManyRequests,
				Data:    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

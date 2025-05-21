package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/gin-gonic/gin"
	responseHelper "github.com/mitra-gh/CarBid/internal/api/response-helper"
	serviceErrors "github.com/mitra-gh/CarBid/pkg/service_errs"
)

func Limiter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusTooManyRequests, responseHelper.Response{
				Code:    -1,
				Message: serviceErrors.ErrTooManyRequests,
				Data:    nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

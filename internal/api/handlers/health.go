package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	responseHelper "github.com/mitra-gh/CarBid/internal/api/response-helper"
)

type HealthHandler struct {}

// HealthCheck godoc
// @Summary Health check
// @Description Health check
// @Accept json
// @Produce json
// @Success 200 {object} responseHelper.Response
// @Failure 400 {object} responseHelper.Response
// @Router /v1/health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, responseHelper.SuccessResponse("OK", "Health check successful"))
}
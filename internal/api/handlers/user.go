package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/api/dto"
	responseHelper "github.com/mitra-gh/CarBid/internal/api/response-helper"
	"github.com/mitra-gh/CarBid/internal/services"
	serviceErrors "github.com/mitra-gh/CarBid/pkg/service_errs"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(cfg *configs.Config) *UserHandler {
	userService := services.NewUserService(cfg)
	return &UserHandler{
		userService: userService,
	}
}


// SendOtp godoc
// @Summary Send OTP
// @Description Send OTP to the user's mobile number
// @Accept json
// @Produce json
// @Param request body dto.GetOtpRequest true "Get OTP Request"
// @Success 200 {object} responseHelper.Response "OTP sent successfully"
// @Failure 400 {object} responseHelper.Response "Invalid request"
// @Failure 500 {object} responseHelper.Response "Failed to send OTP"
// @Router /v1/user/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	var req dto.GetOtpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			responseHelper.ValidationErrorResponse(err, serviceErrors.ErrRequestInvalid ,nil,-1),	
		)
		return
	}

	err := h.userService.SendOtp(req.MobileNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseHelper.ErrorResponse(err, serviceErrors.OtpSendFailed, nil, -1))
		return
	}

	c.JSON(http.StatusOK, responseHelper.SuccessResponse(nil, serviceErrors.OtpSendSuccess))
}


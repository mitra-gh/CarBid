package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,len=11"`
}
package services

import (
	"fmt"

	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/constants"
	"github.com/mitra-gh/CarBid/internal/data/cache"
	"github.com/mitra-gh/CarBid/pkg/logging"
	serviceErrors "github.com/mitra-gh/CarBid/pkg/service_errs"
)

type OtpService struct {
	logger logging.Logger
	cfg    *configs.Config
}

type OtpDto struct {
	Value string `json:"value"`
	Used  bool   `json:"used"`
}


func NewOtpService(cfg *configs.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	return &OtpService{
		logger: logger,
		cfg: cfg,
	}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.OtpKey, mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}

	res,err := cache.Get[OtpDto](key)
	if err == nil && res.Used {
		return serviceErrors.New(serviceErrors.ErrOtpAlreadyExists)
	} else if err == nil {
		return serviceErrors.New(serviceErrors.ErrOtpAlreadyUsed)
	}

	err = cache.Set(key, val, s.cfg.Otp.ExpirationTime)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.OtpKey, mobileNumber)
	res,err := cache.Get[OtpDto](key)

	switch{
		case err != nil:
			return err
		case res.Value != otp:
			return serviceErrors.New(serviceErrors.ErrOtpInvalid)
		case res.Used:
			return serviceErrors.New(serviceErrors.ErrOtpAlreadyUsed)
		case err ==nil && res.Value == otp && !res.Used:
			res.Used = true
			err = cache.Set(key, res, s.cfg.Otp.ExpirationTime)
			if err != nil {
				return err
			}
	}
	return nil
}
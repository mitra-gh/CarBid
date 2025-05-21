package services

import (
	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/common"
	"github.com/mitra-gh/CarBid/internal/data/db"
	"github.com/mitra-gh/CarBid/internal/data/db/sqlc"
	"github.com/mitra-gh/CarBid/pkg/logging"
	serviceErrors "github.com/mitra-gh/CarBid/pkg/service_errs"
)

type UserService struct {
	logger     logging.Logger
	cfg        *configs.Config
	otpService *OtpService
	queries    *sqlc.Queries
}

func NewUserService(cfg *configs.Config) *UserService {
	logger := logging.NewLogger(cfg)
	queries,err := db.GetQueries()
	otpService := NewOtpService(cfg)

	if err != nil {
		logger.Error(logging.Database,logging.StartUp, string(serviceErrors.ErrDBConnectionError), nil)
		return nil
	}
	return &UserService{
		logger: logger,
		cfg: cfg,
		otpService: otpService,
		queries: queries,
	}
}

func (s *UserService) SendOtp(mobileNumber string) error {
	otp := common.OtpGenerator(s.cfg)
	err := s.otpService.SetOtp(mobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
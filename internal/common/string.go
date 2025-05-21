package common

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/mitra-gh/CarBid/configs"
)

func OtpGenerator(cfg *configs.Config) string{
	format := fmt.Sprintf("%s0%dd","%",cfg.Otp.Length)
	tenPow := math.Pow(10, float64(cfg.Otp.Length))
	otp := fmt.Sprintf(format, rand.Intn(int(tenPow)))
	return otp
}
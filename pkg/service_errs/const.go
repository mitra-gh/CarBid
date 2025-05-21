package serviceErrors

type ErrorMassage string

const (
	ErrOtpNotFound      ErrorMassage = "otp not found"
	ErrOtpAlreadyUsed   ErrorMassage = "otp already used"
	ErrOtpExpired       ErrorMassage = "otp expired"
	ErrOtpInvalid       ErrorMassage = "otp invalid"
	ErrOtpAlreadyExists ErrorMassage = "otp already exists"
)

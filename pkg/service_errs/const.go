package serviceErrors

const (

	// otp
	ErrOtpNotFound      = "otp not found"
	ErrOtpAlreadyUsed   = "otp already used"
	ErrOtpExpired       = "otp expired"
	ErrOtpInvalid       = "otp invalid"
	ErrOtpAlreadyExists = "otp already exists"

	// user
	ErrUserNotFound            = "user not found"
	ErrUserAlreadyExists       = "user already exists"
	ErrUserInvalidPassword     = "user invalid password"
	ErrUserInvalidMobileNumber = "user invalid mobile number"
	ErrUserInvalidEmail        = "user invalid email"
	ErrUserInvalidName         = "user invalid name"

	// db
	ErrDBConnectionError = "db connection error"
	ErrDBQueryError      = "db query error"
	ErrDBExecError       = "db exec error"
	ErrDBScanError       = "db scan error"
	ErrDBPrepareError    = "db prepare error"
	ErrDBBeginError      = "db begin error"
	ErrDBCommitError     = "db commit error"
	ErrDBRollbackError   = "db rollback error"

	// request
	ErrRequestInvalid = "request invalid"

	// limiter
	ErrTooManyRequests = "too many requests"
)

const (
	// otp
	OtpSendSuccess = "otp sent successfully"
	OtpSendFailed  = "failed to send otp"

	// user
	UserSendOtpSuccess = "otp sent successfully"
	UserSendOtpFailed  = "failed to send otp"
)
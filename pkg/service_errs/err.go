package serviceErrors

type ServiceError struct {
	massage string
}

func (e *ServiceError) Error() string {
	return e.massage
}

func New(massage string) *ServiceError {
	return &ServiceError{
		massage: massage,
	}
}

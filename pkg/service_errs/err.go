package serviceErrors

type ServiceError struct {
	massage ErrorMassage
}

func (e *ServiceError) Error() string {
	return string(e.massage)
}

func New(massage ErrorMassage) *ServiceError {
	return &ServiceError{
		massage: massage,
	}
}

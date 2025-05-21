package models

type User struct {
	BaseModel
	Username     string `json:"username" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	MobileNumber string `json:"mobile_number" validate:"required,min=10,max=15"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	Enabled      bool   `json:"enabled" validate:"required"`
	UserRoles    *[]UserRole
}

type Role struct {
	BaseModel
	Name      string `json:"name" validate:"required"`
	UserRoles *[]UserRole
}

type UserRole struct {
	BaseModel
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

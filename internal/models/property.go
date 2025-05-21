package models

type PropertyCategory struct {
	BaseModel
	Name       string     `json:"name" validate:"required"`
	Icon       string     `json:"icon" validate:"required"`
	Properties []Property `json:"properties" validate:"required"`
}

type Property struct {
	BaseModel
	Name        string           `json:"name" validate:"required"`
	Icon        string           `json:"icon" validate:"required"`
	Category    PropertyCategory `json:"category" validate:"required"`
	CategoryId  int              `json:"category_id" validate:"required"`
	Description string           `json:"description" validate:"required"`
	DataType    string           `json:"data_type" validate:"required"`
	Unit        string           `json:"unit" validate:"required"`
}
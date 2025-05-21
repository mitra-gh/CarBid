package models

import "time"

type Gearbox struct {
	BaseModel
	Name      string `json:"name" validate:"required"`
	CarModels *[]CarModel
}

type CarType struct {
	BaseModel
	Name      string `json:"name" validate:"required"`
	CarModels *[]CarModel
}

type Company struct {
	BaseModel
	Name      string `json:"name" validate:"required"`
	Country   string `json:"country" validate:"required"`
	CountryID int    `json:"country_id" validate:"required"`
	CarModels *[]CarModel
}

type CarModel struct {
	BaseModel
	Name               string  `json:"name" validate:"required"`
	Company            Company `json:"company" validate:"required"`
	CompanyId          int     `json:"company_id" validate:"required"`
	CarType            CarType `json:"car_type" validate:"required"`
	CarTypeId          int     `json:"car_type_id" validate:"required"`
	Gearbox            Gearbox `json:"gearbox" validate:"required"`
	GearboxId          int     `json:"gearbox_id" validate:"required"`
	CarModelColors     []CarModelColor
	CarModelYears      []CarModelYear
	CarModelProperties []CarModelProperty
	CarModelImages     []CarModelImage
	CarModelComments   []CarModelComment
}

type CarModelColor struct {
	BaseModel
	CarModel   CarModel `json:"car_model" validate:"required"`
	CarModelId int      `json:"car_model_id" validate:"required"`
	Color      Color    `json:"color" validate:"required"`
	ColorId    int      `json:"color_id" validate:"required"`
}

type CarModelYear struct {
	BaseModel
	CarModel               CarModel    `json:"car_model" validate:"required"`
	CarModelId             int         `json:"car_model_id" validate:"required"`
	PersianYear            PersianYear `json:"persian_year" validate:"required"`
	PersianYearId          int         `json:"persian_year_id" validate:"required"`
	CarModelPriceHistories []CarModelPriceHistory
}

type CarModelImage struct {
	BaseModel
	CarModel    CarModel  `json:"car_model" validate:"required"`
	CarModelId  int       `json:"car_model_id" validate:"required"`
	Image       File      `json:"image" validate:"required"`
	ImageId     int       `json:"image_id" validate:"required"`
	IsMainImage bool      `json:"is_main_image" validate:"required"`
}

type CarModelPriceHistory struct {
	BaseModel
	CarModelYear   CarModelYear `json:"car_model_year" validate:"required"`
	CarModelYearId int          `json:"car_model_year_id" validate:"required"`
	Price          float64      `json:"price" validate:"required"`
	PriceAt        time.Time    `json:"price_at" validate:"required"`
}

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `json:"car_model" validate:"required"`
	CarModelId int      `json:"car_model_id" validate:"required"`
	Property   Property `json:"property" validate:"required"`
	PropertyId int      `json:"property_id" validate:"required"`
	Value      string   `json:"value" validate:"required"`
}

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `json:"car_model" validate:"required"`
	CarModelId int      `json:"car_model_id" validate:"required"`
	User       User     `json:"user" validate:"required"`
	UserId     int      `json:"user_id" validate:"required"`
	Message    string   `json:"message" validate:"required"`
}
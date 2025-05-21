package models

import "time"

type Country struct {
	BaseModel
	Name      string `json:"name" validate:"required"`
	Cities    []City
	Companies []Company
}

type City struct {
	BaseModel
	Name      string `json:"name" validate:"required"`
	CountryId int    `json:"country_id" validate:"required"`
	Country   Country `json:"country" validate:"required"`
}

type PersianYear struct {
	BaseModel
	PersianTitle  string    `json:"persian_title" validate:"required"`
	Year          int       `json:"year" validate:"required"`
	StartAt       time.Time `json:"start_at" validate:"required"`
	EndAt         time.Time `json:"end_at" validate:"required"`
	CarModelYears []CarModelYear
}

type Color struct {
	BaseModel
	Name           string `json:"name" validate:"required"`
	HexCode        string `json:"hex_code" validate:"required"`
	CarModelColors []CarModelColor
}

type File struct {
	BaseModel
	Name        string `json:"name" validate:"required"`
	Directory   string `json:"directory" validate:"required"`
	Description string `json:"description" validate:"required"`
	MimeType    string `json:"mime_type" validate:"required"`
}
package models

import "errors"

var (
	ErrorNotFoundCountry = errors.New("Not found country")
)

type Country struct {
	ID int `json:"id"`
	CountryName string `json:"country_name"`
	Address []*Address `json:"address" gorm:"foreignKey:CountryID"`
}
func (Country) TableName() string { return "countries" }

type CreateCountry struct {
	CountryName string `json:"country_name" validate:"required,gte=1,lte=50"`
}

func (CreateCountry) TableName() string { return Country{}.TableName() }

type UpdateCountry struct {
	CountryName string `json:"country_name" validate:"required,gte=1,lte=50"`
}

func (UpdateCountry) TableName() string { return Country{}.TableName() }

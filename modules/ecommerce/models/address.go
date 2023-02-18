package models

type Address struct {
	ID           int     `json:"id"`
	StreetNumber string  `json:"street_number"`
	AddressLine  string  `json:"address_line"`
	City         string  `json:"city"`
	User         []*User `json:"users" gorm:"many2many:user_addresses;"`
	CountryID    int     `json:"country_id"`
}

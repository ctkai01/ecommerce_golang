package models

type Country struct {
	ID int `json:"id"`
	CountryName string `json:"country_name"`
	Address []*Address `json:"address" gorm:"foreignKey:CountryID"`
}

// func (Country) TableName() string {
// 	return "countries"
// }
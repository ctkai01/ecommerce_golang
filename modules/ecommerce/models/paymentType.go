package models

import "time"

type PaymentType struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	User      []*User   `json:"users" gorm:"many2many:user_payment_methods;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

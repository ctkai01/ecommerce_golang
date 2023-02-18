package models

import (
	"ecommerce_shop/common"
	"time"
	// "gorm.io/gorm"
)

type User struct {
	common.SQLModel
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	ShoppingCart []*ShoppingCart `json:"shopping_carts"`
	Address      []*Address      `json:"address" gorm:"many2many:user_addresses;"`
	PaymentType  []*PaymentType  `json:"payment_types" gorm:"many2many:user_payment_methods;"`
}

func (User) TableName() string { return "users" }

type UserCreation struct {
	Id int `json:"-"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,gte=8,lte=20"`
	FullName string `json:"full_name" validate:"required,gte=2,lte=30"`
}

func (UserCreation) TableName() string { return User{}.TableName() }

type UserAddress struct {
	UserID    int       `gorm:"primaryKey"`
	AddressID int       `gorm:"primaryKey"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPaymentMethod struct {
	ID            int       `json:"id" gorm:"column:id;autoIncrement"`
	PaymentTypeID int       `gorm:"primaryKey"`
	UseID         int       `gorm:"primaryKey"`
	Provider      string    `json:"provider"`
	AccountNumber string    `json:"account_number"`
	ExpiryDate    time.Time `json:"expiry_date"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

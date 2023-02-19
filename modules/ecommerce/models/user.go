package models

import (
	"ecommerce_shop/common"
	"time"
	// "gorm.io/gorm"
)

type User struct {
	common.SQLModel
	Email        string          `json:"email"`
	Phone        string          `json:"phone"`
	Password     string          `json:"password"`
	Token        string          `json:"token"`
	FullName     string          `json:"full_name"`
	IsAdmin      bool            `json:"is_admin"`
	ShoppingCart []*ShoppingCart `json:"shopping_carts"`
	Address      []*Address      `json:"address" gorm:"many2many:user_addresses;"`
	PaymentType  []*PaymentType  `json:"payment_types" gorm:"many2many:user_payment_methods;"`
}

func (User) TableName() string { return "users" }

type UserCreation struct {
	Id       int    `json:"-"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,gte=8,lte=20"`
	FullName string `json:"full_name" validate:"required,gte=2,lte=30"`
}

func (UserCreation) TableName() string { return User{}.TableName() }

type UserLogin struct {
	Id       int    `json:"-"`
	Email    string `json:"email" validate:"required,email"`
	Token    string `json:"token"`
	Password string `json:"password" validate:"required,gte=8,lte=20"`
}

func (UserLogin) TableName() string { return User{}.TableName() }

type UserAuth struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (UserAuth) TableName() string { return User{}.TableName() }

type UserUpdate struct {
	FullName string `json:"full_name" validate:"gte=2,lte=30"`
}

func (UserUpdate) TableName() string { return User{}.TableName() }

type UserUpdatePassword struct {
	OldPassword string `json:"old_password" validate:"gte=8,lte=20"`
	Password    string `json:"password" validate:"gte=8,lte=20"`
}

func (UserUpdatePassword) TableName() string { return User{}.TableName() }

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

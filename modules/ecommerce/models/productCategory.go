package models

import (
	"errors"
	"time"
)

var (
	ErrorNotFoundProductCategory = errors.New("Product category not found")
)

type ProductCategory struct {
	ID               int               `json:"id"`
	CategoryName     string               `json:"category_name"`
	CategoryID       *int              `json:"category_id"`
	ChildrenCategories []ProductCategory `json:"children_categories" gorm:"foreignkey:CategoryID;OnDelete:NULL"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

func (ProductCategory) TableName() string { return "product_categories" }

type CreateProductCategory struct {
	ID           int  `json:"-"`
	CategoryName string  `json:"category_name" validate:"required,gte=1,lte=100"`
	CategoryID   *int `json:"category_id"`
}

func (CreateProductCategory) TableName() string { return ProductCategory{}.TableName() }


type UpdateProductCategory struct {
	CategoryName string  `json:"category_name" validate:"omitempty,gte=1,lte=100"`
	CategoryID   *int `json:"category_id"`
}

func (UpdateProductCategory) TableName() string { return ProductCategory{}.TableName() }

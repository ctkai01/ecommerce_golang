package resource

import (
	"ecommerce_shop/modules/ecommerce/models"
	"time"
)

type UserResource struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUserResource(data models.User) UserResource {
	return UserResource{
		ID: data.ID,
		Email: data.Email,
		Phone: data.Phone,
		FullName: data.FullName,
		CreatedAt: *data.CreatedAt,
		UpdatedAt: *data.CreatedAt,
	}
}
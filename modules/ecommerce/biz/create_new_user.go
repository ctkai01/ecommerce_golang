package biz

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
	"errors"
)

// import "context"

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *models.UserCreation) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
} 


func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *models.UserCreation) error {
	return errors.New("sds")

}
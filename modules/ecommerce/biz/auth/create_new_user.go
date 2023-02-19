package biz_auth

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"

	"golang.org/x/crypto/bcrypt"
)

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

	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		return err
	}

	data.Password = string(password)

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}

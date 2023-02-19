package biz_user

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"

	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordStorage interface {
	CurrentUser(c context.Context, id int, data *models.UserAuth) error
	ChangePassword(c context.Context, id int, data *models.UserUpdatePassword) error
}

type changePasswordBiz struct {
	store ChangePasswordStorage
}

func NewChangePasswordBiz(store ChangePasswordStorage) *changePasswordBiz {
	return &changePasswordBiz{
		store: store,
	}
}

func (biz *changePasswordBiz) ChangePassword(c context.Context, id int, data *models.UserUpdatePassword) error {
	var userAuth models.UserAuth

	if err := biz.store.CurrentUser(c, id, &userAuth); err != nil {
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(data.OldPassword)); err != nil {
		return common.ErrorInvalidCredentials
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	if err != nil {
		return err
	}
	data.Password = string(passwordHash)
	biz.store.ChangePassword(c, id, data)
	

	return nil
}
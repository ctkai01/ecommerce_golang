package biz_user

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type CurrentUserStorage interface {
	CurrentUser(c context.Context, id int, data *models.UserAuth) error
}

type currentUserBiz struct {
	store CurrentUserStorage
}

func NewCurrentUserBiz(store CurrentUserStorage) *currentUserBiz {
	return &currentUserBiz{store: store}
}

func (biz *currentUserBiz) CurrentUser(c context.Context, id int, data *models.UserAuth) error {
	
	if err := biz.store.CurrentUser(c, id, data); err != nil {
		return err
	}

	return nil
}



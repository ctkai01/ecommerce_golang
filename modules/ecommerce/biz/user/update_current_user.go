package biz_user

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type UpdateCurrentUserStorage interface {
	UpdateCurrentUser(c context.Context, id int, dataUpdate *models.UserUpdate) error 
}

type updateCurrentUserBiz struct {
	store UpdateCurrentUserStorage
}


func NewUpdateCurrentBiz(store UpdateCurrentUserStorage) *updateCurrentUserBiz {
	return &updateCurrentUserBiz{
		store: store,
	}
}

func (biz *updateCurrentUserBiz) UpdateCurrentUser(c context.Context, id int, dataUpdate *models.UserUpdate) error {
	if err := biz.store.UpdateCurrentUser(c, id, dataUpdate); err != nil {
		return err
	}
	return nil
}



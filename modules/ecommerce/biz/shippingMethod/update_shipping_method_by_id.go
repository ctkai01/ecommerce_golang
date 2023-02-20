package biz_shipping_method

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type UpdateShippingMethodStorage interface {
	GetShippingMethod(ctx context.Context, id int) error
	UpdateShippingMethod(ctx context.Context, id int, dataUpdate *models.UpdateShippingMethod ) error
}

type updateShippingMethodBiz struct {
	store UpdateShippingMethodStorage
}

func NewUpdateShippingMethodBiz(store UpdateShippingMethodStorage) *updateShippingMethodBiz {
	return &updateShippingMethodBiz{store: store}
} 

func (biz *updateShippingMethodBiz) UpdateShippingMethodByID(ctx context.Context, id int, dataUpdate *models.UpdateShippingMethod) error {

	if err := biz.store.GetShippingMethod(ctx, id); err != nil {
		return err
	}
	
	if err := biz.store.UpdateShippingMethod(ctx, id, dataUpdate); err != nil {
		return err
	}
	return nil
}
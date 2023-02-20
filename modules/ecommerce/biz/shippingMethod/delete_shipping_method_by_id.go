package biz_shipping_method

import (
	"context"
)

type DeleteShippingMethodStorage interface {
	GetShippingMethod(ctx context.Context, id int) error
	DeleteShippingMethod(ctx context.Context, id int) error
}

type deleteShippingMethodBiz struct {
	store DeleteShippingMethodStorage
}

func NewDeleteShippingMethodBiz(store DeleteShippingMethodStorage) *deleteShippingMethodBiz {
	return &deleteShippingMethodBiz{store: store}
} 

func (biz *deleteShippingMethodBiz) DeleteShippingMethodByID(ctx context.Context, id int) error {

	if err := biz.store.GetShippingMethod(ctx, id); err != nil {
		return err
	}
	
	if err := biz.store.DeleteShippingMethod(ctx, id); err != nil {
		return err
	}
	return nil
}
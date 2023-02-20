package biz_shipping_method

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"
)

type ListShippingMethodStorage interface {
	ListShippingMethod(ctx context.Context, paging *common.Paging) ([]models.ShippingMethod, error)
}

type listShippingMethodBiz struct {
	store ListShippingMethodStorage
}

func ListShippingMethodBiz(store ListShippingMethodStorage) *listShippingMethodBiz {
	return &listShippingMethodBiz{store: store}
} 

func (biz *listShippingMethodBiz) ListShippingMethod(ctx context.Context, paging *common.Paging) ([]models.ShippingMethod, error) {
	data, err := biz.store.ListShippingMethod(ctx, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
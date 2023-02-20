package biz_shipping_method
import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type CreateShippingMethodStorage interface {
	CreateShippingMethod(ctx context.Context, data *models.CreateShippingMethod) error
}

type creatShippingMethodBiz struct {
	store CreateShippingMethodStorage
}

func NewCreateShippingMethodBiz(store CreateShippingMethodStorage) *creatShippingMethodBiz {
	return &creatShippingMethodBiz{store: store}
} 

func  (biz *creatShippingMethodBiz) CreateShippingMethod(ctx context.Context, data *models.CreateShippingMethod) error {
	if err := biz.store.CreateShippingMethod(ctx, data); err != nil {
		return err
	}

	return nil
}
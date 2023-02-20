
package biz_payment_type
import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type CreatePaymentTypeStorage interface {
	CreatePaymentType(ctx context.Context, data *models.CreatePaymentType) error
}

type creatPaymentTypeBiz struct {
	store CreatePaymentTypeStorage
}

func NewCreatePaymentTypeBiz(store CreatePaymentTypeStorage) *creatPaymentTypeBiz {
	return &creatPaymentTypeBiz{store: store}
} 

func  (biz *creatPaymentTypeBiz) CreatePaymentType(ctx context.Context, data *models.CreatePaymentType) error {
	if err := biz.store.CreatePaymentType(ctx, data); err != nil {
		return err
	}

	return nil
}
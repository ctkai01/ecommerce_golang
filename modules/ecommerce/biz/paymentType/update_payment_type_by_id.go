package biz_payment_type

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type UpdatePaymentTypeStorage interface {
	GetPaymentType(ctx context.Context, id int) error
	UpdatePaymentType(ctx context.Context, id int, dataUpdate *models.UpdatePaymentType ) error
}

type updatePaymentTypeBiz struct {
	store UpdatePaymentTypeStorage
}

func NewUpdatePaymentTypeBiz(store UpdatePaymentTypeStorage) *updatePaymentTypeBiz {
	return &updatePaymentTypeBiz{store: store}
} 

func (biz *updatePaymentTypeBiz) UpdatePaymentTypeByID(ctx context.Context, id int, dataUpdate *models.UpdatePaymentType) error {

	if err := biz.store.GetPaymentType(ctx, id); err != nil {
		return err
	}
	
	if err := biz.store.UpdatePaymentType(ctx, id, dataUpdate); err != nil {
		return err
	}
	return nil
}
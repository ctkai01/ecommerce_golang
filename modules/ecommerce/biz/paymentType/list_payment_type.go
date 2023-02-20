package biz_payment_type

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"
)

type ListPaymentTypeStorage interface {
	ListPaymentType(ctx context.Context, paging *common.Paging) ([]models.PaymentType, error)
}

type listPaymentTypeBiz struct {
	store ListPaymentTypeStorage
}

func ListPaymentTypeBiz(store ListPaymentTypeStorage) *listPaymentTypeBiz {
	return &listPaymentTypeBiz{store: store}
} 

func (biz *listPaymentTypeBiz) ListPaymentType(ctx context.Context, paging *common.Paging) ([]models.PaymentType, error) {
	data, err := biz.store.ListPaymentType(ctx, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
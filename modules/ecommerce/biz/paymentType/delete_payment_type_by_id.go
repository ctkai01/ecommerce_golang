package biz_payment_type

import (
	"context"
)

type DeletePaymentTypeStorage interface {
	GetPaymentType(ctx context.Context, id int) error
	DeletePaymentType(ctx context.Context, id int) error
}

type deletePaymentTypeBiz struct {
	store DeletePaymentTypeStorage
}

func NewDeletePaymentTypeBiz(store DeletePaymentTypeStorage) *deletePaymentTypeBiz {
	return &deletePaymentTypeBiz{store: store}
} 

func (biz *deletePaymentTypeBiz) DeletePaymentTypeByID(ctx context.Context, id int) error {

	if err := biz.store.GetPaymentType(ctx, id); err != nil {
		return err
	}
	
	if err := biz.store.DeletePaymentType(ctx, id); err != nil {
		return err
	}
	return nil
}
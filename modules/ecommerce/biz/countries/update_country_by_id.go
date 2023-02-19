package biz_countries

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type UpdateCountryStorage interface {
	GetCountry(ctx context.Context, id int) error
	UpdateCountry(ctx context.Context, id int, dataUpdate *models.UpdateCountry ) error
}

type updateCountryBiz struct {
	store UpdateCountryStorage
}

func NewUpdateCountryBiz(store UpdateCountryStorage) *updateCountryBiz {
	return &updateCountryBiz{store: store}
} 

func (biz *updateCountryBiz) UpdateCountryByID(ctx context.Context, id int, dataUpdate *models.UpdateCountry) error {

	if err := biz.store.GetCountry(ctx, id); err != nil {
		return err
	}
	
	if err := biz.store.UpdateCountry(ctx, id, dataUpdate); err != nil {
		return err
	}
	return nil
}
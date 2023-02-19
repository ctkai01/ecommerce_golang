package biz_countries

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type CreateCountryStorage interface {
	CreateCountry(ctx context.Context, data *models.CreateCountry) error
}

type createCountryBiz struct {
	store CreateCountryStorage
}

func NewCreateCountryBiz(store CreateCountryStorage) *createCountryBiz {
	return &createCountryBiz{store: store}
} 

func  (biz *createCountryBiz) CreateCountry(ctx context.Context, data *models.CreateCountry) error {
	if err := biz.store.CreateCountry(ctx, data); err != nil {
		return err
	}

	return nil
}
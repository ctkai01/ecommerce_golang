package biz_countries

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"
)

type ListCountryStorage interface {
	ListCountry(ctx context.Context, paging *common.Paging) ([]models.Country, error)
}

type listCountryBiz struct {
	store ListCountryStorage
}

func ListCountryBiz(store ListCountryStorage) *listCountryBiz {
	return &listCountryBiz{store: store}
} 

func (biz *listCountryBiz) ListCountry(ctx context.Context, paging *common.Paging) ([]models.Country, error) {
	data, err := biz.store.ListCountry(ctx, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
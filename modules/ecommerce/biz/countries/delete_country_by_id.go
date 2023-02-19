package biz_countries

import (
	"context"
)

type DeleteCountryStorage interface {
	GetCountry(ctx context.Context, id int) error
	DeleteCountry(ctx context.Context, id int) error
}

type deleteCountryBiz struct {
	store DeleteCountryStorage
}

func NewDeleteCountryBiz(store DeleteCountryStorage) *deleteCountryBiz {
	return &deleteCountryBiz{store: store}
} 

func (biz *deleteCountryBiz) DeleteCountryByID(ctx context.Context, id int) error {

	if err := biz.store.GetCountry(ctx, id); err != nil {
		return err
	}
	
	if err := biz.store.DeleteCountry(ctx, id); err != nil {
		return err
	}
	return nil
}
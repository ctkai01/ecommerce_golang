package biz_product_category

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type GetProductCategoryStorage interface {
	FindProductCategory(ctx context.Context, id int) (*models.ProductCategory, error)
}

type getProductCategoryBiz struct {
	store GetProductCategoryStorage
}

func NewGetProductCategoryBiz(store GetProductCategoryStorage) *getProductCategoryBiz {
	return &getProductCategoryBiz{store: store}
} 

func (biz *getProductCategoryBiz) GetProductCategory(ctx context.Context, id int) (*models.ProductCategory, error) {
	data, err := biz.store.FindProductCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
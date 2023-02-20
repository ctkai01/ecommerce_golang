package biz_product_category

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"
)

type ListProductCategoryStorage interface {
	ListProductCategory(ctx context.Context, paging *common.Paging) ([]models.ProductCategory, error)
}

type listProductCategoryBiz struct {
	store ListProductCategoryStorage
}

func NewListProductCategoryBiz(store ListProductCategoryStorage) *listProductCategoryBiz {
	return &listProductCategoryBiz{store: store}
} 

func (biz *listProductCategoryBiz) ListProductCategory(ctx context.Context, paging *common.Paging) ([]models.ProductCategory, error) {
	data, err := biz.store.ListProductCategory(ctx, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
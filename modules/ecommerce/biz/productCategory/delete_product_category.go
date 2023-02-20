package biz_product_category

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type DeleteProductCategoryStorage interface {
	FindProductCategory(ctx context.Context, id int) (*models.ProductCategory, error)
	DeleteProductCategory(ctx context.Context, id int) error
}

type deleteProductCategoryBiz struct {
	store DeleteProductCategoryStorage
}

func NewDeleteProductCategoryBiz(store DeleteProductCategoryStorage) *deleteProductCategoryBiz {
	return &deleteProductCategoryBiz{store: store}
} 

func (biz *deleteProductCategoryBiz) DeleteProductCategoryByID(ctx context.Context, id int) error {

	_, err := biz.store.FindProductCategory(ctx, id)

	if err != nil {
		return err
	}
	
	if err := biz.store.DeleteProductCategory(ctx, id); err != nil {
		return err
	}
	return nil
}
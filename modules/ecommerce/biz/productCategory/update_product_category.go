package biz_product_category

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
	"fmt"
)

type UpdateProductCategoryStorage interface {
	FindProductCategory(ctx context.Context, id int) (*models.ProductCategory, error)
	UpdateProductCategory(ctx context.Context, id int, dataUpdate *models.UpdateProductCategory) error
}

type updateProductCategoryBiz struct {
	store UpdateProductCategoryStorage
}

func NewUpdateProductCategoryBiz(store UpdateProductCategoryStorage) *updateProductCategoryBiz {
	return &updateProductCategoryBiz{store: store}
}

func (biz *updateProductCategoryBiz) UpdateProductCategoryByID(ctx context.Context, id int, dataUpdate *models.UpdateProductCategory) error {
	_, err := biz.store.FindProductCategory(ctx, id)
	if err != nil {
		return err
	}

	fmt.Println(dataUpdate.CategoryName == "")
	fmt.Println(dataUpdate)
	if dataUpdate.CategoryID != nil{
		if *dataUpdate.CategoryID != 0 {
			_, err = biz.store.FindProductCategory(ctx, *dataUpdate.CategoryID)
			if err != nil {
				return err
			}
		}
	}
	
	if err := biz.store.UpdateProductCategory(ctx, id, dataUpdate); err != nil {
		return err
	}
	return nil
}

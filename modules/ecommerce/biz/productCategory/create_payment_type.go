
package biz_product_category
import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

type CreateProductCategoryStorage interface {
	FindProductCategory(ctx context.Context, id int) (*models.ProductCategory, error)
	CreateProductCategory(ctx context.Context, data *models.CreateProductCategory) error
}

type creatProductCategoryBiz struct {
	store CreateProductCategoryStorage
}

func NewCreateProductCategory(store CreateProductCategoryStorage) *creatProductCategoryBiz {
	return &creatProductCategoryBiz{store: store}
} 

func  (biz *creatProductCategoryBiz) CreateProductCategory(ctx context.Context, data *models.CreateProductCategory) error {
	if data.CategoryID != nil {
		_, err := biz.store.FindProductCategory(ctx, *data.CategoryID)
		if err != nil {
			return err
		}
	}
	
	if err := biz.store.CreateProductCategory(ctx, data); err != nil {
		return err
	}

	return nil
}


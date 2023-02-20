package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)
func (s *sqlStore) FindProductCategory(ctx context.Context, id int) (*models.ProductCategory, error) {
	var product_category models.ProductCategory 
	if err := s.db.Preload("ChildrenCategories").First(&product_category, id).Error; err != nil {
		return nil, models.ErrorNotFoundProductCategory
	}

	return &product_category, nil
}
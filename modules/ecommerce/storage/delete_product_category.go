package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) DeleteProductCategory(ctx context.Context, id int) error {
	if err := s.db.Delete(&models.ProductCategory{}, id).Error; err != nil {
		return err
	}

	return nil
}

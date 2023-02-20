package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) GetShippingMethod(ctx context.Context, id int) error {
	var result *models.ShippingMethod
	
	if err := s.db.First(&result, id).Error; err != nil {
		return models.ErrorNotFoundShippingMethod
	}

	return nil
}
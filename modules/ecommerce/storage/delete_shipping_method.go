package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) DeleteShippingMethod(ctx context.Context, id int) error {
	if err := s.db.Delete(&models.ShippingMethod{}, id).Error; err != nil {
		return err
	}

	return nil
}

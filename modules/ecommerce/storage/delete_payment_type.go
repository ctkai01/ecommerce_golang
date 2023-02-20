package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) DeletePaymentType(ctx context.Context, id int) error {
	if err := s.db.Delete(&models.CreatePaymentType{}, id).Error; err != nil {
		return err
	}

	return nil
}

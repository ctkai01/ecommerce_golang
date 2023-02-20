package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) GetPaymentType(ctx context.Context, id int) error {
	var result *models.PaymentType
	
	if err := s.db.First(&result, id).Error; err != nil {
		return models.ErrorNotFoundPaymentType
	}

	return nil
}
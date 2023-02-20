package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) CreatePaymentType(ctx context.Context, data *models.CreatePaymentType) error {
	if err := s.db.Table(data.TableName()).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
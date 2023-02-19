package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) CreateCountry(ctx context.Context, data *models.CreateCountry) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
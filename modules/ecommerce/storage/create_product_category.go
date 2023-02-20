package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) CreateProductCategory(ctx context.Context, data *models.CreateProductCategory) error {
	if err := s.db.Table(data.TableName()).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
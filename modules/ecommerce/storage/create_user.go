package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)


func (s *sqlStore) CreateUser(ctx context.Context, data *models.UserCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

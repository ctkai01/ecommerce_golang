package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) Logout(ctx context.Context, id int) error {
	if err := s.db.Model(&models.User{}).Where("id = ?", id).Update("token", nil).Error; err != nil {
		return err
	}
	return nil
}
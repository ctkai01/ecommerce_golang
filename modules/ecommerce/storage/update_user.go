package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)


func (s *sqlStore) UpdateCurrentUser(ctx context.Context, id int, data *models.UserUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

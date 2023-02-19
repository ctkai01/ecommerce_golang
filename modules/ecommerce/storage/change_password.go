package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)


func (s *sqlStore) ChangePassword(c context.Context, id int, data *models.UserUpdatePassword) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).Update("password", data.Password).Error; err != nil {
		return err
	}

	return nil
}

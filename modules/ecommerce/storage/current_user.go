package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) CurrentUser(c context.Context, id int, data *models.UserAuth) error {
	data.ID = id
	
	if err := s.db.First(&data).Error; err != nil {
		return err
	}
	
	return nil
}
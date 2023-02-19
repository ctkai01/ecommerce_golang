package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) UpdateCountry(ctx context.Context, id int, dataUpdate *models.UpdateCountry) error {
	if err := s.db.Table(dataUpdate.TableName()).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}

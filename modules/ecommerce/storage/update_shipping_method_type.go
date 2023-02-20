package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) UpdateShippingMethod(ctx context.Context, id int, dataUpdate *models.UpdateShippingMethod) error {
	if err := s.db.Table(dataUpdate.TableName()).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}

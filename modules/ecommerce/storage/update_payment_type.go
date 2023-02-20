package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) UpdatePaymentType(ctx context.Context, id int, dataUpdate *models.UpdatePaymentType) error {
	if err := s.db.Table(dataUpdate.TableName()).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}

package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) UpdateProductCategory(ctx context.Context, id int, dataUpdate *models.UpdateProductCategory) error {
	updates := map[string]interface{}{
		"category_id":   dataUpdate.CategoryID,
	}

	if dataUpdate.CategoryName != "" {
		updates["category_name"] = dataUpdate.CategoryName
	}

	if dataUpdate.CategoryID != nil{
		if *dataUpdate.CategoryID == 0 {
			updates["category_id"] = nil
		}
	}

	if err := s.db.Table(dataUpdate.TableName()).Where("id = ?", id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

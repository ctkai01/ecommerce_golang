package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) GetCountry(ctx context.Context, id int) error {
	var result *models.Country
	
	if err := s.db.First(&result, id).Error; err != nil {
		return models.ErrorNotFoundCountry
	}

	return nil
}
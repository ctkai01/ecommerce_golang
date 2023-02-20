package storage

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) ListShippingMethod(
	ctx context.Context, paging *common.Paging,
) ([]models.ShippingMethod, error) {
	var result []models.ShippingMethod
	
	if err := s.db.Table(models.ShippingMethod{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := s.db.Order("id desc").
	Offset((paging.Page - 1) * paging.Limit).
	Limit(paging.Limit).
	Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
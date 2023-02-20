package storage

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) ListPaymentType(
	ctx context.Context, paging *common.Paging,
) ([]models.PaymentType, error) {
	var result []models.PaymentType
	
	if err := s.db.Table(models.PaymentType{}.TableName()).Count(&paging.Total).Error; err != nil {
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
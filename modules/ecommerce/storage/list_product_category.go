package storage

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) ListProductCategory(
	ctx context.Context, paging *common.Paging,
) ([]models.ProductCategory, error) {
	var result []models.ProductCategory
	
	if err := s.db.Table(models.ProductCategory{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := s.db.Debug().Where("category_id IS NULL") .Order("id desc").
	Offset((paging.Page - 1) * paging.Limit).
	Limit(paging.Limit).
	Preload("ChildrenCategories").
	Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
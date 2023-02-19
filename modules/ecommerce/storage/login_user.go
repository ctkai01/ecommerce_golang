package storage

import (
	"context"
	"ecommerce_shop/modules/ecommerce/models"
)

func (s *sqlStore) 	LoginUser(ctx context.Context, data *models.UserLogin) error {
	if err := s.db.Table(data.TableName()).Where("email = ?", data.Email).First(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) SetToken(ctx context.Context, data *models.UserLogin, token string) error {
	if err := s.db.Table(data.TableName()).Where("email = ?", data.Email).Update("token", token).Error; err != nil {
		return err
	}
	return nil
}
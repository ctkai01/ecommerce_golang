package biz_auth

import (
	"context"
	"ecommerce_shop/common"
	"ecommerce_shop/configs"
	"ecommerce_shop/modules/ecommerce/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserStorage interface {
	LoginUser(ctx context.Context, data *models.UserLogin) error
	SetToken(ctx context.Context, data *models.UserLogin, token string) error
}

type loginUserBiz struct {
	store LoginUserStorage
}

func NewLoginUserBiz(store LoginUserStorage) *loginUserBiz {
	return &loginUserBiz {
		store: store,
	}
}


func (biz *loginUserBiz) LoginUser(ctx context.Context, data *models.UserLogin) error {
	password_input := data.Password	
	if err := biz.store.LoginUser(ctx, data); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(password_input)); err != nil {
		return common.ErrorInvalidCredentials
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"sub": strconv.Itoa(data.Id),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
 	token, err := claims.SignedString([]byte(configs.EnvConfigs.Secret))

	if err != nil {
		return err 
	}

	data.Token = token
	if err := biz.store.SetToken(ctx, data, token); err != nil {
		return err
	}
	
	return nil
}
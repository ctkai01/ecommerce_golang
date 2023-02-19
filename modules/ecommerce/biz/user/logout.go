package biz_user

import "context"

type LogOutStorage interface {
	Logout(ctx context.Context, id int) error
}

type logOutBiz struct {
	store LogOutStorage
}

func NewLogoutBiz(store LogOutStorage) *logOutBiz {
	return &logOutBiz {
		store: store,
	}
}

func (biz *logOutBiz) Logout(ctx context.Context, id int) error {
	biz.store.Logout(ctx, id)
	
	return nil
}

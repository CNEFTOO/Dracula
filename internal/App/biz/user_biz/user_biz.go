package userbiz

import (
	"context"

	"github.com/seaung/Dracula/internal/App/store"
)

type UserBiz interface {
	UserLogin(ctx context.Context, username, password string) (bool, error)
}

type userBiz struct {
	ds store.Store
}

func NewUserBiz(ds store.Store) *userBiz {
	return &userBiz{ds: ds}
}

var _ UserBiz = (*userBiz)(nil)

func (u *userBiz) UserLogin(ctx context.Context, username, password string) (bool, error) {
	return false, nil
}

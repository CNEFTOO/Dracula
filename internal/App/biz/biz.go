package biz

import (
	userbiz "github.com/seaung/Dracula/internal/App/biz/user_biz"
	wsdmbiz "github.com/seaung/Dracula/internal/App/biz/wsdm_biz"
	"github.com/seaung/Dracula/internal/App/store"
)

type Biz interface {
	User() userbiz.UserBiz
	Wsdms() wsdmbiz.WsdmBiz
}

var _ Biz = (*biz)(nil)

type biz struct {
	ds store.Store
}

func NewBiz(ds store.Store) *biz {
	return &biz{
		ds: ds,
	}
}

func (b *biz) User() userbiz.UserBiz {
	return userbiz.NewUserBiz(b.ds)
}

func (b *biz) Wsdms() wsdmbiz.WsdmBiz {
	return wsdmbiz.NewWsdmBiz(b.ds)
}

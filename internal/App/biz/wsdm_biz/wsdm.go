package wsdmbiz

import (
	"context"

	"github.com/seaung/Dracula/internal/App/store"
)

type WsdmBiz interface {
	List(ctx context.Context) error
}

type wsdmBiz struct {
	ds store.Store
}

func NewWsdmBiz(ds store.Store) *wsdmBiz {
	return &wsdmBiz{ds: ds}
}

var _ WsdmBiz = (*wsdmBiz)(nil)

func (wb *wsdmBiz) List(ctx context.Context) error {
	return nil
}

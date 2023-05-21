package biz

import (
	"github.com/seaung/hystrix/internal/hystrix/biz/user"
	"github.com/seaung/hystrix/internal/hystrix/store"
)

type ISBiz interface {
	Users() user.UserBiz
}

type biz struct {
	ds store.IsStore
}

var _ ISBiz = (*biz)(nil)

func NewBiz(ds store.IsStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

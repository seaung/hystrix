package users

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/hystrix/internal/hystrix/biz"
	"github.com/seaung/hystrix/internal/hystrix/store"
)

type UserControllers struct {
	b biz.ISBiz
}

func New(ds store.IsStore) *UserControllers {
	return &UserControllers{b: biz.NewBiz(ds)}
}

func (u *UserControllers) Login(c *gin.Context) {
}

func (u *UserControllers) UpdatePassword(c *gin.Context) {
}

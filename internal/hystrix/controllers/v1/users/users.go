package users

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/hystrix/internal/hystrix/biz"
	"github.com/seaung/hystrix/internal/hystrix/store"
	"github.com/seaung/hystrix/internal/pkg/core"
	"github.com/seaung/hystrix/internal/pkg/logs"
	v1 "github.com/seaung/hystrix/pkg/api/hystrix/v1"
)

type UserControllers struct {
	b biz.ISBiz
}

func New(ds store.IsStore) *UserControllers {
	return &UserControllers{b: biz.NewBiz(ds)}
}

func (u *UserControllers) Login(c *gin.Context) {
	logs.C(c).Infow("Login function called")
	var r v1.LoginFormRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, nil, gin.H{"code": 10012, "httpCode": 200, "data": "参数校验失败"})
		return
	}

	resp, err := u.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, nil, gin.H{"code": 10013, "httpCode": 200, "data": "nil"})
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (u *UserControllers) UpdatePassword(c *gin.Context) {
}

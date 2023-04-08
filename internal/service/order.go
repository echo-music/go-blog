package service

import (
	"errors"
	"github.com/echo-music/go-blog/internal/model"
	"github.com/echo-music/go-blog/pkg/api/goblog"
	"github.com/echo-music/go-blog/pkg/g"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/gin-gonic/gin"
)

type userSrv struct {
}

var User userSrv

func (a *userSrv) List(c *gin.Context, arg goblog.UserListArg) (users goblog.UserListRet, err error) {
	panic("草泥马")
	err = g.DB().Model(&model.User{}).Scan(&users.List).Error
	return
}

func (a *userSrv) Create(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))
	if err != nil {
		return
	}

}

func (a *userSrv) Update(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))
	if err != nil {
		return
	}

}

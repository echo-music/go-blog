package service

import (
	"errors"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/echo-music/go-blog/pkg/model"
	"github.com/echo-music/go-blog/pkg/store/mysql"
	"github.com/gin-gonic/gin"
)

type userSrv struct {
}

var User userSrv

func (a *userSrv) List(c *gin.Context, arg model.UserListArg) (users model.UserListRet, err error) {
	err = mysql.DB().Debug().Model(&model.User{}).Scan(&users.List).Error
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

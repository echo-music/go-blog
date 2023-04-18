package service

import (
	"errors"
	"fmt"
	"github.com/echo-music/go-blog/internal/model"
	"github.com/echo-music/go-blog/pkg/api/goblog"
	"github.com/echo-music/go-blog/pkg/g"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/singleflight"
)

type userSrv struct {
}

var single singleflight.Group
var User userSrv
var res goblog.UserListRet

func (a *userSrv) List(c *gin.Context, arg goblog.UserListArg) (users goblog.UserListRet, err error) {
	if len(res.List) > 0 {
		users = res
		return users, nil
	}

	_, err = single.Do("cache", func() (interface{}, error) {
		fmt.Println("查询db")
		err = g.DB().Model(&model.User{}).Scan(&res.List).Error
		return nil, err
	})
	if err != nil {
		return
	}

	users = res

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

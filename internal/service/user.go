package service

import (
	"fmt"
	"github.com/echo-music/go-blog/internal/model"
	"github.com/echo-music/go-blog/pkg/api/blog"
	"github.com/echo-music/go-blog/pkg/g"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/singleflight"
	"github.com/pkg/errors"
	"os"
)

type userSrv struct {
}

var single singleflight.Group
var User userSrv
var res blog.UserListRet

func (a *userSrv) List(c *gin.Context, arg blog.UserListArg) (users blog.UserListRet, err error) {
	if len(res.List) > 0 {
		users = res
		return users, nil
	}
	if _, err = os.Open("a.txt"); err != nil {
		return res, gerror.New(errors.WithStack(err))
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

}

func (a *userSrv) Update(c *gin.Context) {

}

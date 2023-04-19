package api

import (
	"github.com/echo-music/go-blog/internal/service"
	"github.com/echo-music/go-blog/pkg/api/blog"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
)

type userApi struct {
}

var User userApi

func (a *userApi) Router(r *gin.Engine) {
	user := r.Group("/users")
	user.Use()
	{
		user.GET("/", User.List)
		user.POST("/", User.Create)
		user.PUT("/", User.Update)

	}
}

// List
// @Summary 用户列表
// @Tags 用户管理
// @param   _ query blog.UserListArg _ "_"
// @success	200 {object} response.Result{data=blog.UserListRet} "_"
// @Router /users [get]
func (a *userApi) List(c *gin.Context) {

	var arg blog.UserListArg
	if err := c.ShouldBindQuery(&arg); err != nil {
		c.Error(err)
		return
	}

	list, err := service.User.List(c, arg)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, "用户列表大傻逼", list)
}

func (a *userApi) Create(c *gin.Context) {

	response.Success(c, "用户创建", nil)
}

func (a *userApi) Update(c *gin.Context) {

	response.Success(c, "修改用户", nil)
}

package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-blog/internal/service"
	"go-blog/pkg/gerror"
	"go-blog/pkg/response"
)

type orderApi struct {
}

var Order orderApi

func (a *orderApi) Router(r *gin.Engine) {
	order := r.Group("/orders")
	order.Use()
	{
		order.GET("/", Order.List)
		order.POST("/", Order.Create)
		order.PUT("/", Order.Update)

	}
}

func (a *orderApi) List(c *gin.Context) {

	list := service.Order.List(c)

	response.Success(c, "订单列表", list)
}

func (a *orderApi) Create(c *gin.Context) {

	response.Success(c, "订单创建", nil)
}

func (a *orderApi) Update(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))

	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, "修改订单", nil)
}

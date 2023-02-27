package api

import (
	"errors"
	"github.com/echo-music/go-blog/internal/service"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/echo-music/go-blog/pkg/model"
	"github.com/echo-music/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
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

// List
// PingExample godoc
// @Summary 订单列表
// @Schemes
// @Description do ping
// @Tags 订单管理
// @Accept json
// @Produce json
// @param   _ query model.OrderListArg _ "_"
// @success	200 {object} response.Result{data=model.OrderListRet} "_"
// @Router /orders [get]
func (a *orderApi) List(c *gin.Context) {

	arg := model.OrderListArg{}
	list, err := service.Order.List(c, arg)
	if err != nil {
		c.Error(err)
		return
	}

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

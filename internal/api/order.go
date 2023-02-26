package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-blog/internal/service"
	"go-blog/pkg/library/gerror"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"message": "订单列表",
		"data":    list,
	})
}

func (a *orderApi) Create(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "订单创建",
	})
}

func (a *orderApi) Update(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改订单",
	})
}

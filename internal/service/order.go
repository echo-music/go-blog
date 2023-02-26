package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-blog/pkg/library/gerror"
	"net/http"
)

type orderSrv struct {
}

var Order orderSrv

func (a *orderSrv) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "订单列表",
	})
}

func (a *orderSrv) Create(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "订单创建",
	})
}

func (a *orderSrv) Update(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改订单",
	})
}

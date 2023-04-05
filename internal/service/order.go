package service

import (
	"errors"
	"github.com/echo-music/go-blog/pkg/db"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/echo-music/go-blog/pkg/model"
	"github.com/gin-gonic/gin"
)

type orderSrv struct {
}

var Order orderSrv

func (a *orderSrv) List(c *gin.Context, arg model.OrderListArg) (orders model.OrderListRet, err error) {
	err = db.DB().Model(&model.Order{}).Scan(&orders.List).Error
	return
}

func (a *orderSrv) Create(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))
	if err != nil {
		return
	}

}

func (a *orderSrv) Update(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))
	if err != nil {
		return
	}

}

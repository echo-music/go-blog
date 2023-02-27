package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-blog/pkg/db"
	"go-blog/pkg/gerror"
	"go-blog/pkg/model"
	"net/http"
)

type orderSrv struct {
}

var Order orderSrv

func (a *orderSrv) List(c *gin.Context) (orders []*model.Order) {

	db.DB().Model(&model.Order{}).Scan(&orders)
	return
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

package service

import (
	"errors"
	"github.com/echo-music/go-blog/pkg/db"
	"github.com/echo-music/go-blog/pkg/gerror"
	"github.com/echo-music/go-blog/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type orderSrv struct {
}

var Order orderSrv

func (a *orderSrv) List(c *gin.Context, arg model.OrderListArg) (orders model.OrderListRet, err error) {

	err = db.DB().Model(&model.Order{}).Scan(&orders.List).Error
	tr := otel.Tracer("order-list")
	_, span := tr.Start(c.Request.Context(), "list", oteltrace.WithAttributes(attribute.String("id", "100")))

	span.SetAttributes(attribute.Int("age", 10))
	span.SetAttributes(attribute.String("name", "张三"))

	defer span.End()

	otelzap.Ctx(c.Request.Context()).Error("order-list", zap.String("name", "你这次不上百哦"))

	a.Create(c)
	return
}

func (a *orderSrv) Create(c *gin.Context) {
	err := gerror.New(errors.New("订单号不能为空"))

	tr := otel.Tracer("order-Create")
	_, span := tr.Start(c.Request.Context(), "Create", oteltrace.WithAttributes(attribute.String("id", "100")))

	span.SetAttributes(attribute.Int("age", 10))
	span.SetAttributes(attribute.String("name", "kkk"))
	span.RecordError(err)
	defer span.End()

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

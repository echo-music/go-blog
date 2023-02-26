package model

type Order struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (Order) TableName() string {
	return "order"
}

type OrderCreateArg struct {
}

type OrderCreateRet struct {
}

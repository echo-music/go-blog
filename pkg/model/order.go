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

type OrderListArg struct {
}

type OrderListRet struct {
	List []OrderListRow
}

type OrderListRow struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

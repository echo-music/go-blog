package model

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (User) TableName() string {
	return "user"
}

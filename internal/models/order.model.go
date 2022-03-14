package models

import "time"

var OrderTname string = "order"

type Order struct {
	Order_id       int       `json:"order_id"`
	Order_datatime time.Time `json:"order_datatime" validate:"required"`
	Order_status   int       `json:"order_status" gorm:"size:64" validate:"required"`
	Order_user_id  int       `json:"order_user_id" gorm:"size:64" validate:"required"`
	Order_addr     int       `json:"order_addr" gorm:"size:64" validate:"required"`
}

func (order *Order) TableName() string {
	return "order"
}

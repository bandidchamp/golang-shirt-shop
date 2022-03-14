package models

var OrderItemTname string = "order_item"

type OrderItem struct {
	Order_item_id         int `json:"order_item_id"`
	Order_item_quantity   int `json:"order_item_quantity" gorm:"size:64" validate:"required"`
	Order_item_product_id int `json:"order_item_product_id" gorm:"size:64" validate:"required"`
	Order_id              int `json:"order_id" gorm:"size:64" validate:"required"`
}

func (OrderItem *OrderItem) TableName() string {
	return "order_item"
}

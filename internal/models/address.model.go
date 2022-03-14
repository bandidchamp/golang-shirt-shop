package models

var AddressTname string = "user_address"

type Address struct {
	Addr_id       int    `json:"addr_id" gorm:"size:64"`
	Addr_addresss string `json:"addr_addresss" validate:"required"`
	Addr_user_id  int    `json:"addr_user_id" gorm:"size:64" validate:"required"`
}

func (addr *Address) TableName() string {
	return "user_address"
}

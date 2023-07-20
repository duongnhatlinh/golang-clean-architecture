package model

import "food_delivery/common"

const EntityName = "Order"

type Order struct {
	common.SQLModel `json:",inline"`
	UserId          int     `json:"user_id" gorm:"column:user_id"`
	TotalPrice      float32 `json:"total_price" gorm:"column:total_price"`
	ShipperId       int     `json:"shipper_id" gorm:"column:shipper_id"`
}

func (Order) TableName() string {
	return "orders"
}

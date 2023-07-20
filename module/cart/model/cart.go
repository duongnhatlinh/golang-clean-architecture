package model

import (
	"food_delivery/module/food/model"
	"time"
)

const EntityName = "Cart"

type Cart struct {
	UserId    int         `json:"user_id" gorm:"column:user_id"`
	FoodId    int         `json:"food_id" gorm:"column:food_id"`
	Quantity  int         `json:"quantity" gorm:"column:quantity"`
	Status    int         `json:"status" gorm:"column:status;default:1"`
	Food      *model.Food `json:"food" gorm:"preload:false"`
	CreatedAt *time.Time  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (Cart) TableName() string {
	return "carts"
}

type GetParamCart struct {
	FoodId   string `form:"food_id"`
	Quantity int    `form:"quantity"`
}

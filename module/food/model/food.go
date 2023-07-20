package model

import (
	"food_delivery/common"
	catModel "food_delivery/module/category/model"
	resModel "food_delivery/module/restaurant/model"
)

const EntityName = "Food"

type Food struct {
	common.SQLModel `json:",inline"`
	RestaurantId    int                 `json:"-" gorm:"column:restaurant_id"`
	Restaurant      resModel.Restaurant `json:"restaurant" gorm:"foreignKey:RestaurantId;preload:false"`
	CategoryId      int                 `json:"-" gorm:"category_id"`
	Category        catModel.Category   `json:"category" gorm:"foreignKey:CategoryId;preload:false"`
	Name            string              `json:"name" gorm:"column:name"`
	Description     string              `json:"description" gorm:"column:description"`
	Price           float32             `json:"price" gorm:"column:price"`
	Image           *common.Image       `json:"image" gorm:"column:images"`
}

func (Food) TableName() string {
	return "foods"
}

type FoodForeignKey struct {
	RestaurantId string `json:"restaurant_id" form:"restaurant_id"`
	CategoryId   int    `json:"category_id" form:"category_id"`
}

func (f *Food) Mask(isAdminOrUser bool) {
	f.GenUID(common.DbTypeFood)
}

type UpdateFood struct {
	Name        *string       `json:"name" gorm:"column:name"`
	Description *string       `json:"description" gorm:"column:description"`
	Price       *float32      `json:"price" gorm:"column:price"`
	Image       *common.Image `json:"image" gorm:"column:images"`
}

func (UpdateFood) TableName() string {
	return Food{}.TableName()
}

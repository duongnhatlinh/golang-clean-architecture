package model

import "food_delivery/common"

const EntityName = "Category"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name"`
	Description     string        `json:"description" gorm:"column:description"`
	Icon            *common.Image `json:"icon" gorm:"icon"`
}

func (Category) TableName() string {
	return "categories"
}

func (f *Category) Mask(isAdminOrUser bool) {
	f.GenUID(common.DbTypeFood)
}

type UpdateCategory struct {
	Name        *string       `json:"name" gorm:"column:name"`
	Description *string       `json:"description" gorm:"column:description"`
	Icon        *common.Image `json:"icon" gorm:"icon"`
}

func (UpdateCategory) TableName() string {
	return Category{}.TableName()
}

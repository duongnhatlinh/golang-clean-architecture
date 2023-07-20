package model

import "food_delivery/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	OwnerId         int                `json:"-" gorm:"column:owner_id"`
	Name            string             `json:"name" gorm:"column:name"`
	Address         string             `json:"address" gorm:"column:addr"`
	User            *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerId;preload:false"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover"`
	LikedCount      int                `json:"liked_count" gorm:"column:liked_count"` // computed field
}

type UpdateRestaurant struct {
	Name    *string        `json:"name" gorm:"column:name"`
	Address *string        `json:"address" gorm:"column:addr"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo"`
	Cover   *common.Images `json:"cover" gorm:"column:cover"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (UpdateRestaurant) TableName() string {
	return Restaurant{}.TableName()
}

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

	if user := data.User; user != nil {
		user.Mask(isAdminOrOwner)
	}
}

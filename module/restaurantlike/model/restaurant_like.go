package model

import (
	"food_delivery/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Restaurant_like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserId       int                `json:"user_id" gorm:"column:user_id"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"created_at"`
}

func (Restaurant_like) TableName() string {
	return "restaurant_likes"
}

func (r *Restaurant_like) GetRestaurantId() int {
	return r.RestaurantId
}

func (r *Restaurant_like) GetUserId() int {
	return r.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot like this restaurant",
		"ErrCannotLikeRestaurant",
	)
}

func ErrCannotDisLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot dis like this restaurant",
		"ErrCannotDisLikeRestaurant",
	)
}

package model

import (
	"food_delivery/common"
	"time"
)

const EntityName = "UserLikeFood"

type Food_like struct {
	FoodId    int                `json:"food_id" gorm:"column:food_id"`
	UserId    int                `json:"user_id" gorm:"column:user_id"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false"`
	CreatedAt *time.Time         `json:"created_at" gorm:"created_at"`
}

func (Food_like) TableName() string {
	return "food_likes"
}

func (r *Food_like) GetFoodId() int {
	return r.FoodId
}

func (r *Food_like) GetUserId() int {
	return r.UserId
}

func ErrCannotLikeFood(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot like this food",
		"ErrCannotLikeFood",
	)
}

func ErrCannotDisLikeFood(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot dislike this food",
		"ErrCannotDisLikeFood",
	)
}

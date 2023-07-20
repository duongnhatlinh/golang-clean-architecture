package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *mysqlStorage) UpdateRestaurant(ctx context.Context, id int, data *model.UpdateRestaurant) error {
	db := s.DB

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

func (s *mysqlStorage) IncreaseLikeCountRestaurant(ctx context.Context, id int) error {
	db := s.DB

	if err := db.Table(model.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

func (s *mysqlStorage) DescendLikeCountRestaurant(ctx context.Context, id int) error {
	db := s.DB

	if err := db.Table(model.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

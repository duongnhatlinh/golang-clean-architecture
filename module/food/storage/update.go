package storage

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/food/model"
	"gorm.io/gorm"
)

func (s *mysqlStorage) UpdateFood(
	ctx context.Context,
	condition map[string]interface{},
	data *model.UpdateFood,
) error {
	db := s.db

	if err := db.Where(condition).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

func (s *mysqlStorage) IncreaseLikeCountFood(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.Food{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

func (s *mysqlStorage) DescendLikeCountFood(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.Food{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}

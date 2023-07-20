package business

import (
	"context"
	"food_delivery/module/order/model"
)

type CreateOrderRepo interface {
	CreateOrder(ctx context.Context, data *model.Order) error
}

type createOrderBiz struct {
	repo CreateOrderRepo
}

func NewCreateOrderBiz(repo CreateOrderRepo) *createOrderBiz {
	return &createOrderBiz{repo: repo}
}

func (biz *createOrderBiz) CreateNewOrder(ctx context.Context, data *model.Order) error {
	if err := biz.repo.CreateOrder(ctx, data); err != nil {
		return err
	}
	return nil
}

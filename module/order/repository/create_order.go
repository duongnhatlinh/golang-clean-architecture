package repository

import (
	"context"
	"food_delivery/common"
	modelCart "food_delivery/module/cart/model"
	"food_delivery/module/order/model"
)

type CreateOrderStorage interface {
	Create(ctx context.Context, data *model.Order) error
}

type cartStorage interface {
	List(ctx context.Context, userId int, moreKeys ...string) ([]modelCart.Cart, error)
}

type createOrderRepo struct {
	store     CreateOrderStorage
	cartStore cartStorage
}

func NewCreateOrderRepo(store CreateOrderStorage, cartStore cartStorage) *createOrderRepo {
	return &createOrderRepo{
		store:     store,
		cartStore: cartStore,
	}
}

func (repo *createOrderRepo) CreateOrder(ctx context.Context, data *model.Order) error {
	cart, err := repo.cartStore.List(ctx, data.UserId, "Food")

	if err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	var sumPrice float32
	for i := range cart {
		sumPrice += float32(cart[i].Quantity) * cart[i].Food.Price
	}
	data.TotalPrice = sumPrice

	if err := repo.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}

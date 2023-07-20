package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.User) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterBiz(store RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{store: store, hasher: hasher}
}

func (biz *registerBiz) RegisterUser(ctx context.Context, data *model.User) error {
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return model.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}

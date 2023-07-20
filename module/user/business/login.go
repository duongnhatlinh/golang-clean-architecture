package business

import (
	"context"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/tokenprovider"
	"food_delivery/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*model.User, error)
}

type loginBusiness struct {
	store         LoginStorage
	appCtx        appctx.AppContext
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(
	store LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *loginBusiness {
	return &loginBusiness{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// login
// 1. Find user by email
// 2. Hash pass from input and compare pass in db
// 3. Provider: issue JWT token for client
// 4. Return token

func (l *loginBusiness) LoginUser(ctx context.Context, data *model.UserLogin, moreKeys ...string) (*tokenprovider.Token, error) {
	user, err := l.store.FindUser(ctx, map[string]interface{}{"email": data.Email}, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	passHashed := l.hasher.Hash(data.Password + user.Salt)

	if passHashed != user.Password {
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	token, err := l.tokenProvider.Generate(payload, l.expiry)

	if err != nil {
		return nil, err
	}

	return token, nil
}

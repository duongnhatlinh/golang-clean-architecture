package appctx

import (
	"food_delivery/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnect() *gorm.DB
	GetSecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	pb        pubsub.Pubsub
}

func NewAppCtx(db *gorm.DB, secretKey string, pb pubsub.Pubsub) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		pb:        pb,
	}
}

func (ctx *appCtx) GetMainDBConnect() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.pb
}

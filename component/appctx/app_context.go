package appctx

import (
	"thucidol/pubsub"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	SecretKey() string
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	ps        pubsub.PubSub
}

func NewAppCtx(db *gorm.DB, secretKey string, ps pubsub.PubSub) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, ps: ps}
}

func (ctx *appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.ps
}

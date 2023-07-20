package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/cart/business"
	"food_delivery/module/cart/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRemoveFoodFromCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewRemoveFoodBiz(store)

		if err := biz.RemoveFoodFromCart(ctx.Request.Context(), user.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/cart/business"
	"food_delivery/module/cart/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetListFoodsFromCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewGetListFoodsBiz(store)

		data, err := biz.GetListFoodsFromCart(ctx.Request.Context(), user.GetUserId(), "Food")
		if err != nil {
			panic(err)
		}

		for _, item := range data {

			item.Food.Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRespond(data, nil, nil))
	}
}

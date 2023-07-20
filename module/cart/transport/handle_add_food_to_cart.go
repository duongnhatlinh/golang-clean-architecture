package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/cart/business"
	"food_delivery/module/cart/model"
	"food_delivery/module/cart/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleAddFoodToCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var para model.GetParamCart
		if err := ctx.ShouldBindQuery(&para); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uid, err := common.FromBase58(para.FoodId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data model.Cart
		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = user.GetUserId()
		data.FoodId = int(uid.GetLocalID())
		data.Quantity = para.Quantity

		if data.Quantity < 1 {
			data.Quantity = 1
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewAddFoodBiz(store)

		if err := biz.AddFoodToCart(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

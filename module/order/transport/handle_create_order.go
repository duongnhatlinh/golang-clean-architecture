package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	storageCart "food_delivery/module/cart/storage"
	"food_delivery/module/order/business"
	"food_delivery/module/order/model"
	"food_delivery/module/order/repository"
	"food_delivery/module/order/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleCreateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type shipper struct {
			ShipperId int `form:"shipper_id"`
		}

		var ship shipper

		if err := ctx.ShouldBindQuery(&ship); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data model.Order
		data.UserId = user.GetUserId()
		data.ShipperId = ship.ShipperId
		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		cartStore := storageCart.NewMysqlStorage(appCtx.GetMainDBConnect())
		repo := repository.NewCreateOrderRepo(store, cartStore)
		biz := business.NewCreateOrderBiz(repo)

		if err := biz.CreateNewOrder(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

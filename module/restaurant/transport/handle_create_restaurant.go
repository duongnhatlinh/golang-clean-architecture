package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurant/business"
	"food_delivery/module/restaurant/model"
	"food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleCreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.Restaurant

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		data.OwnerId = user.GetUserId()

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewCreateRestaurantBiz(store)

		if err := biz.CreateNewRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))

	}
}

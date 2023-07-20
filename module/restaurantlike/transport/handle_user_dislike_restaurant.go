package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurantlike/business"
	"food_delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idRes, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		restaurantId := int(idRes.GetLocalID())
		userId := user.GetUserId()

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewUserDislikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.UserDislikeRestaurant(ctx.Request.Context(), userId, restaurantId); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

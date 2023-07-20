package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurantlike/business"
	"food_delivery/module/restaurantlike/model"
	"food_delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idRes, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		data := model.Restaurant_like{
			RestaurantId: int(idRes.GetLocalID()),
			UserId:       user.GetUserId(),
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.UserLikeRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

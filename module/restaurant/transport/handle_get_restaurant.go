package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurant/business"
	"food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		uid, err := common.FromBase58(ctx.Param("id"))

		//id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(data))
	}
}

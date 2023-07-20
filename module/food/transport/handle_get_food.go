package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/business"
	"food_delivery/module/food/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewGetFoodBiz(store)

		data, err := biz.GetFood(ctx.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(data))
	}
}

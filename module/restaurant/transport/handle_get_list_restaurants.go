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

func HandleGetListRestaurants(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter model.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewGetListRestaurantBiz(store)

		data, err := biz.ListRestaurants(ctx.Request.Context(), &filter, &paging, "User")

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)

			if i == len(data)-1 {
				paging.NextCursor = data[i].FakeId.String()
			}
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRespond(data, paging, filter))
	}
}

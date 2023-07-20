package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/business"
	"food_delivery/module/food/model"
	"food_delivery/module/food/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetListFoods(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter model.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		paging.Fulfill()

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewGetListFoodsBiz(store)

		data, err := biz.GetListFoods(ctx.Request.Context(), nil, &paging, &filter, "Restaurant")

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)

			if i == len(data)-1 {
				paging.NextCursor = data[i].FakeId.String()
			}

			data[i].Restaurant.Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRespond(data, filter, paging))
	}
}

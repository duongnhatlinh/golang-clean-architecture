package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/category/business"
	"food_delivery/module/category/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetListCategories(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var paging common.Paging

		paging.Fulfill()

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewGetListCategoriesBiz(store)

		data, err := biz.GetListCategories(ctx.Request.Context(), nil, &paging)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)

			if i == len(data)-1 {
				paging.NextCursor = data[i].FakeId.String()
			}
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRespond(data, nil, paging))
	}
}

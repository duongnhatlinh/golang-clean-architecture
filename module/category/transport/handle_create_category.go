package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/category/business"
	"food_delivery/module/category/model"
	"food_delivery/module/category/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleCreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data model.Category
		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewCreateCategoryBiz(store)

		if err := biz.CreateNewCategory(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

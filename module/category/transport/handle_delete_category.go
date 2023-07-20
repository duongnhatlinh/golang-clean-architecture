package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/category/business"
	"food_delivery/module/category/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleDeleteCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewDeleteCategoryBiz(store)

		if err := biz.RemoveCategory(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

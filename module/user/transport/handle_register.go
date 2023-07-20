package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/module/user/business"
	"food_delivery/module/user/model"
	"food_delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRegister(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.User

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		hasher := hasher.NewMd5Hash()

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewRegisterBiz(store, hasher)

		if err := biz.RegisterUser(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(data.FakeId.String()))
	}
}

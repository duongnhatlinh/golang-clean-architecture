package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/component/tokenprovider"
	"food_delivery/module/user/business"
	"food_delivery/module/user/model"
	"food_delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleLogin(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data model.UserLogin

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		hasher := hasher.NewMd5Hash()
		tokenProvider := tokenprovider.NewTokenJWTProvider(appCtx.GetSecretKey())
		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewLoginBusiness(store, tokenProvider, hasher, 60*60*24)

		token, err := biz.LoginUser(ctx.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(token))
	}
}

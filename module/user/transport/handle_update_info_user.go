package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/user/business"
	"food_delivery/module/user/model"
	"food_delivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUpdateInfoUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data model.UserUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewUpdateInfoUserBiz(store)

		if err := biz.UpdateInfoUser(ctx.Request.Context(), user.GetUserId(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

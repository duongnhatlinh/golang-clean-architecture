package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/foodlike/business"
	"food_delivery/module/foodlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUserDislikeFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idRes, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		foodId := int(idRes.GetLocalID())
		userId := user.GetUserId()

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewUserDislikeFoodBiz(store, appCtx.GetPubSub())

		if err := biz.UserDislikeFood(ctx.Request.Context(), userId, foodId); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

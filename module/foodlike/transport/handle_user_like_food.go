package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/foodlike/business"
	"food_delivery/module/foodlike/model"
	"food_delivery/module/foodlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUserLikeFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idRes, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		data := model.Food_like{
			FoodId: int(idRes.GetLocalID()),
			UserId: user.GetUserId(),
		}

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewUserLikeFoodBiz(store, appCtx.GetPubSub())

		if err := biz.UserLikeFood(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

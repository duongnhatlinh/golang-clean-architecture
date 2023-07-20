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

func HandleListUserLikeFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		filter := model.Filter{
			FoodId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewListUserLikeFoodBiz(store)

		users, err := biz.ListUserLikeFood(ctx.Request.Context(), nil, &paging, &filter, "User")

		if err != nil {
			panic(err)
		}

		for i := range users {
			users[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRespond(users, paging, filter))
	}
}

package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/restaurantlike/business"
	"food_delivery/module/restaurantlike/model"
	"food_delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleListUserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		filter := model.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewListUserLikeRestaurantBiz(store)

		users, err := biz.ListUserLikeRestaurant(ctx.Request.Context(), nil, &paging, &filter, "User")

		if err != nil {
			panic(err)
		}

		for i := range users {
			users[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRespond(users, paging, filter))
	}
}

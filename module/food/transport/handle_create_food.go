package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/module/food/business"
	"food_delivery/module/food/model"
	"food_delivery/module/food/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleCreateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data model.Food
		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var val model.FoodForeignKey
		if err := ctx.ShouldBindQuery(&val); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//convert int for restaurant
		uid, err := common.FromBase58(val.RestaurantId)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.CategoryId = val.CategoryId
		data.RestaurantId = int(uid.GetLocalID())

		store := storage.NewMysqlStorage(appCtx.GetMainDBConnect())
		biz := business.NewCreateFoodBiz(store)

		if err := biz.CreateNewFood(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(true))
	}
}

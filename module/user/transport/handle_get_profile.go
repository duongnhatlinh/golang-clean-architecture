package transport

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetProfile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := ctx.MustGet(common.CurrentUser).(common.Requester)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespond(data))
	}
}

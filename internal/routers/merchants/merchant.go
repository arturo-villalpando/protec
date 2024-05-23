package merchants

import (
	"github.com/gin-gonic/gin"
)

func (r *MerchantRouter) MerchantRouter(router *gin.Engine) {
	router.POST("/api/merchant/", r.controller.CreateMerchant)
	router.GET("/api/merchant/:id", r.controller.ReadMerchant)
	router.PATCH("/api/merchant/:id", r.controller.UpdateMerchant)
}

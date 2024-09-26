package routes

import (
	"github.com/gin-gonic/gin"
	"payment/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/make-payment", handler.MakePayment)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"ticket-registry/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/get-tickets", handler.GetTicketsForGivenTypeAndQuantity)
}

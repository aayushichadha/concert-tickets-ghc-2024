package routes

import (
	"book-tickets/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/book-tickets", handler.BookTickets)
}

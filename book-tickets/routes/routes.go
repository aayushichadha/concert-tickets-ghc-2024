package routes

import (
	"book-tickets/handler"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers the application routes
func SetupRoutes(r *gin.Engine, bookingHandler *handler.BookingHandler) {
	// Register the BookTickets route
	r.POST("/book-tickets", bookingHandler.BookTickets)
}

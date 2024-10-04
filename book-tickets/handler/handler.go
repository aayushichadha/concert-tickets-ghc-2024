package handler

import (
	"book-tickets/models"
	"book-tickets/service"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookingHandler handles the booking requests
type BookingHandler struct {
	BookingService *service.BookingService
	DB             *gorm.DB
}

// BookTickets is the Gin handler that processes ticket booking requests
func (h *BookingHandler) BookTickets(c *gin.Context) {
	// Step 1: Bind the request body to BookTicketsRequest struct
	var bookTicketsRequest models.BookTicketsRequest
	if err := c.ShouldBindJSON(&bookTicketsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Step 2: Call the BookTickets function in the service layer
	bookedTickets, err := h.BookingService.BookTickets(h.DB, &bookTicketsRequest)
	if err != nil {
		log.Println("Error booking tickets:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Step 3: Return the booked tickets in the response
	c.JSON(http.StatusOK, bookedTickets)
}

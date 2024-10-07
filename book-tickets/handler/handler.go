package handler

import (
	"book-tickets/models"
	"book-tickets/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

// BookingHandler handles the booking requests
type BookingHandler struct {
	BookingService *service.BookingService
	DB             *gorm.DB
	Logger         *logrus.Logger
}

// BookTickets is the Gin handler that processes ticket booking requests
func (h *BookingHandler) BookTickets(c *gin.Context) {
	var bookTicketsRequest models.BookTicketsRequest

	// Create a reusable log entry with endpoint and method fields
	logEntry := h.Logger.WithFields(logrus.Fields{
		"endpoint": "/book-tickets",
		"method":   c.Request.Method,
	})

	// Log the incoming request
	logEntry.Info("Received booking request")

	// Parse the request payload
	if err := c.ShouldBindJSON(&bookTicketsRequest); err != nil {
		logEntry.WithField("error", err.Error()).Warn("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Add user-specific context to the log entry
	logEntry = logEntry.WithFields(logrus.Fields{
		"user_id":    bookTicketsRequest.User.Id,
		"ticketType": bookTicketsRequest.Tickets.Type,
		"quantity":   bookTicketsRequest.Tickets.Quantity,
	})

	// Log processing of the booking request
	logEntry.Info("Processing booking request")

	// Call the service layer to handle the booking
	bookedTickets, err := h.BookingService.BookTickets(h.DB, h.Logger, &bookTicketsRequest)
	if err != nil {
		logEntry.WithField("error", err.Error()).Error("Error booking tickets")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log successful ticket booking
	logEntry.Info("Tickets successfully booked")

	// Return the successful response
	c.JSON(http.StatusOK, bookedTickets)
}

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

	h.Logger.WithFields(logrus.Fields{
		"endpoint": "/book-tickets",
		"method":   c.Request.Method,
	}).Info("Received booking request")

	if err := c.ShouldBindJSON(&bookTicketsRequest); err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	h.Logger.WithFields(logrus.Fields{
		"user_id":    bookTicketsRequest.User.Id,
		"ticketType": bookTicketsRequest.Tickets.Type,
		"quantity":   bookTicketsRequest.Tickets.Quantity,
	}).Info("Processing booking request")

	bookedTickets, err := h.BookingService.BookTickets(h.DB, h.Logger, &bookTicketsRequest)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"user_id": bookTicketsRequest.User.Id,
			"error":   err.Error(),
		}).Error("Error booking tickets")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookedTickets)
}

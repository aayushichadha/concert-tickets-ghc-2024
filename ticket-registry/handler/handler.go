package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ticket-registry/logging"
	"ticket-registry/models"
	"ticket-registry/service"
	"time"
)

func GetTicketsForGivenTypeAndQuantity(c *gin.Context) {

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var getTicketsRequest *models.GetTicketsRequest
	if err := c.ShouldBindJSON(&getTicketsRequest); err != nil {
		log.Printf("Error while parsing order data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logging.Logger.WithFields(logrus.Fields{
		"ticket_type": getTicketsRequest.TicketType,
		"quantity":    getTicketsRequest.Quantity,
	}).Info("Logging - Received GetTicketsForGivenTypeAndQuantity request")

	// Log the request for tracking
	log.Printf("Received GetTicketsForGivenTypeAndQuantity request: Type=%s, Quantity=%d", getTicketsRequest.TicketType, getTicketsRequest.Quantity)

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.GetTicketsForGivenTypeAndQuantity(db.(*gorm.DB), getTicketsRequest)
	if err != nil {
		log.Printf("Error fetching tickets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func ShowTickets(c *gin.Context) {

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Log the request for tracking
	log.Printf("Received ShowTickets request")

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.ShowTickets(db.(*gorm.DB))
	if err != nil {
		log.Printf("Error fetching tickets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

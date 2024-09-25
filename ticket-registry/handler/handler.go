package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
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

	// Log the request for tracking
	log.Printf("Received GetTickets request: Type=%s, Quantity=%d", getTicketsRequest.TicketType, getTicketsRequest.Quantity)

	db, _ := c.Get("db")
	// Retrieve the database object from the context
	// db, exists := c.Get("db")
	// if !exists {
	//     log.Printf("Database connection not found in context")
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	//     return
	// }

	// Call the service layer to handle the order placement
	tickets, err := service.GetTicketsForGivenTypeAndQuantity(db.(*gorm.DB), getTicketsRequest)
	if err != nil {
		log.Printf("Error fetching tickets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"ticket-registry/models"
	"ticket-registry/service"
	"time"
)

func GetTicketsForGivenTypeAndQuantity(c *gin.Context) {

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticketType := c.Query("ticketType")
	quantityStr := c.Query("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity value"})
		return
	}

	// Log the request for tracking
	log.Printf("Received GetTicketsForGivenTypeAndQuantity request: Type=%s, Quantity=%d", ticketType, quantity)

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.GetTicketsForGivenTypeAndQuantity(db.(*gorm.DB), &models.GetTicketsRequest{
		TicketType: ticketType,
		Quantity:   quantity,
	})
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

package handler

import (
	"book-tickets/models"
	"book-tickets/service"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func BookTickets(c *gin.Context) {
	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var bookTicketsRequest *models.BookTicketsRequest
	if err := c.ShouldBindJSON(&bookTicketsRequest); err != nil {
		log.Printf("Error while parsing order data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the request for tracking

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.BookTickets(db.(*gorm.DB), bookTicketsRequest)
	if err != nil {
		log.Printf("Error fetching tickets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

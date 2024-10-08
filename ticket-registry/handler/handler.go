package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for structured logs
	logger.SetLevel(logrus.InfoLevel)
	logEntry := logger.WithFields(logrus.Fields{
		"endpoint": "/get-tickets",
		"method":   c.Request.Method,
	})

	// Log the request for tracking
	logEntry.Info("Received get-tickets request")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		logEntry.WithField("error", err.Error()).Error("Invalid quantity value payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity value"})
		return
	}

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.GetTicketsForGivenTypeAndQuantity(db.(*gorm.DB), &models.GetTicketsRequest{
		TicketType: ticketType,
		Quantity:   quantity,
	}, logger)
	if err != nil {
		logEntry.WithField("error", err.Error()).Error("Error processing get-tickets")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func ShowTickets(c *gin.Context) {

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for structured logs
	logger.SetLevel(logrus.InfoLevel)
	logEntry := logger.WithFields(logrus.Fields{
		"endpoint": "/show-tickets",
		"method":   c.Request.Method,
	})

	// Log the request for tracking
	logEntry.Info("Received show-tickets request")

	db, _ := c.Get("db")

	// Call the service layer to handle the order placement
	tickets, err := service.ShowTickets(db.(*gorm.DB), logger)
	if err != nil {
		logEntry.WithField("error", err.Error()).Error("Error processing show-tickets")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

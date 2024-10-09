package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"payment/models"
	"payment/service"
)

func MakePayment(c *gin.Context) {
	var paymentRequest models.MakePaymentRequest
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{}) // Use JSON format for structured logs
	logger.SetLevel(logrus.InfoLevel)
	logEntry := logger.WithFields(logrus.Fields{
		"endpoint": "/make-payment",
		"method":   c.Request.Method,
	})

	// Log the incoming request
	logEntry.Info("Received payment request")

	// Bind the request body to the MakePaymentRequest struct
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		logEntry.WithField("error", err.Error()).Warn("Invalid payment request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response, err := service.ProcessPayment(paymentRequest, logger)
	if err != nil {
		// Log the error with context
		logEntry.WithField("error", err.Error()).Error("Error processing payment")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the successful response in JSON format
	c.JSON(http.StatusOK, response)
}

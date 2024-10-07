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

	// Add payment-specific context to the log entry
	logEntry = logEntry.WithFields(logrus.Fields{
		"user_id":     paymentRequest.UserID,
		"amount":      paymentRequest.Amount,
		"paymentType": paymentRequest.PaymentMethod.Type,
	})

	// Log the payment request details
	logEntry.Info("Processing payment request")

	// Process the payment using the service layer
	response, err := service.ProcessPayment(paymentRequest, logger)
	if err != nil {
		// Log the error with context
		logEntry.WithField("error", err.Error()).Error("Error processing payment")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the successful payment
	logEntry.Info("Payment processed successfully")

	// Return the successful response in JSON format
	c.JSON(http.StatusOK, response)
}

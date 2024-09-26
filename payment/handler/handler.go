package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/models"
	"payment/service"
)

func MakePayment(c *gin.Context) {

	var paymentRequest models.MakePaymentRequest

	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		// If binding fails, return an error response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Process the payment using the service layer
	response, err := service.ProcessPayment(paymentRequest)
	if err != nil {
		// If there is a business logic error, return the error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the successful response in JSON format
	c.JSON(http.StatusOK, response)
}

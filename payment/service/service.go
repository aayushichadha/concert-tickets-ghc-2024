package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"payment/models"
)

// ProcessPayment processes the payment request and returns a response
func ProcessPayment(request models.MakePaymentRequest, logger *logrus.Logger) (models.MakePaymentResponse, error) {

	// Simulate basic validation
	if request.UserID == "" || request.PaymentMethod.Authorization == "" {
		logger.Warn("Validation failed for payment request")
		return models.MakePaymentResponse{}, errors.New("invalid payment request")
	}

	// Simulate successful payment processing
	response := models.MakePaymentResponse{
		PaymentId: "98765xyz",
		Status:    "Success",
	}

	// Log successful payment
	logger.WithField("paymentId", response.PaymentId).Info("Payment processed successfully")

	return response, nil
}

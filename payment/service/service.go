package service

import (
	"errors"
	"payment/models"
)

// ProcessPayment processes the payment request and returns a response
func ProcessPayment(request models.MakePaymentRequest) (models.MakePaymentResponse, error) {
	// Simulate basic validation
	if request.User.Id == "" || request.PaymentMethod.Authorization == "" {
		return models.MakePaymentResponse{}, errors.New("invalid payment request")
	}

	// Simulate successful payment processing
	response := models.MakePaymentResponse{
		PaymentId: "98765xyz",
		Status:    "Success",
	}

	return response, nil
}

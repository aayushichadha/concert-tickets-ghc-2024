package gateways

import (
	"book-tickets/config"
	"book-tickets/models"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PaymentGateway defines the interface for interacting with the payment service
type PaymentGateway interface {
	MakePayment(request *models.MakePaymentRequest) error
}

// PaymentGatewayImpl is the concrete implementation of the PaymentGateway interface
type PaymentGatewayImpl struct {
	paymentServiceURL string
	makePaymentRoute  string
	Logger            *logrus.Logger // Adding Logger for structured logging
}

// NewPaymentGateway creates a new instance of PaymentGatewayImpl
func NewPaymentGateway(logger *logrus.Logger) (PaymentGateway, error) {
	paymentServiceURL, err := config.ReadServiceConfig("payment")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error reading payment service config")
		return nil, err
	}

	makePaymentRoute, err := config.ReadAPIConfig("make-payment")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error reading make-payment API config")
		return nil, err
	}

	return &PaymentGatewayImpl{
		paymentServiceURL: paymentServiceURL,
		makePaymentRoute:  makePaymentRoute,
		Logger:            logger, // Pass the logger to the implementation
	}, nil
}

// MakePayment processes the payment by interacting with the payment service
func (p *PaymentGatewayImpl) MakePayment(request *models.MakePaymentRequest) error {
	// Log the start of the payment process
	p.Logger.WithFields(logrus.Fields{
		"user_id":     request.UserID,
		"amount":      request.Amount,
		"paymentType": request.PaymentMethod.Type,
	}).Info("Initiating payment request")

	// Marshal the payment request into JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		p.Logger.WithFields(logrus.Fields{
			"user_id": request.UserID,
			"error":   err.Error(),
		}).Error("Error marshalling payment request")
		return err
	}

	// Build the full URL for the make-payment API
	makePaymentURL := p.paymentServiceURL + p.makePaymentRoute

	// Make the HTTP POST request
	resp, err := http.Post(makePaymentURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		p.Logger.WithFields(logrus.Fields{
			"user_id": request.UserID,
			"error":   err.Error(),
			"url":     makePaymentURL,
		}).Error("Error making POST request to payment service")
		return err
	}
	defer resp.Body.Close()

	// Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			p.Logger.WithFields(logrus.Fields{
				"user_id": request.UserID,
				"error":   err.Error(),
				"url":     makePaymentURL,
			}).Error("Error reading payment service response")
			return err
		}
		p.Logger.WithFields(logrus.Fields{
			"user_id":    request.UserID,
			"statusCode": resp.StatusCode,
			"response":   string(body),
		}).Error("Payment service returned non-200 status")
		return errors.New(string(body))
	}

	// Log successful payment
	p.Logger.WithFields(logrus.Fields{
		"user_id": request.UserID,
		"amount":  request.Amount,
	}).Info("Payment processed successfully")

	return nil
}

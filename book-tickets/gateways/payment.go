package gateways

import (
	"book-tickets/config"
	"book-tickets/models"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// PaymentGateway defines the interface for interacting with the payment service
type PaymentGateway interface {
	MakePayment(request *models.MakePaymentRequest) error
}

// PaymentGatewayImpl is the concrete implementation of the PaymentGateway interface
type PaymentGatewayImpl struct {
	paymentServiceURL string
	makePaymentRoute  string
}

// NewPaymentGateway creates a new instance of PaymentGatewayImpl
func NewPaymentGateway() (PaymentGateway, error) {
	paymentServiceURL, err := config.ReadServiceConfig("payment")
	if err != nil {
		return nil, err
	}

	makePaymentRoute, err := config.ReadAPIConfig("make-payment")
	if err != nil {
		return nil, err
	}

	return &PaymentGatewayImpl{
		paymentServiceURL: paymentServiceURL,
		makePaymentRoute:  makePaymentRoute,
	}, nil
}

func (p *PaymentGatewayImpl) MakePayment(request *models.MakePaymentRequest) error {
	// Marshal the payment request into JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// Build the full URL for the make payment API
	makePaymentURL := p.paymentServiceURL + p.makePaymentRoute

	// Make the HTTP POST request
	resp, err := http.Post(makePaymentURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	return nil
}

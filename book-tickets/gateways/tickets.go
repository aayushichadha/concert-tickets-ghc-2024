package gateways

import (
	"book-tickets/config"
	"book-tickets/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// TicketRegistryGateway interface for interacting with the ticket registry
type TicketRegistryGateway interface {
	GetTicketsForGivenTypeAndQuantity(request *models.GetTicketsRequest) ([]models.Ticket, error)
}

// TicketRegistryGatewayImpl is the concrete implementation of TicketRegistryGateway
type TicketRegistryGatewayImpl struct {
	ticketRegistryServiceURL string
	getTicketsRoute          string
	Logger                   *logrus.Logger // Add Logger for structured logging
}

// NewTicketRegistryGateway creates a new instance of TicketRegistryGatewayImpl
func NewTicketRegistryGateway(logger *logrus.Logger) (TicketRegistryGateway, error) {
	// Fetch service URL and API route from config
	ticketRegistryServiceURL, err := config.ReadServiceConfig("ticketRegistry")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error reading ticket registry service config")
		return nil, err
	}

	getTicketsRoute, err := config.ReadAPIConfig("get-tickets")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error reading get-tickets API config")
		return nil, err
	}

	return &TicketRegistryGatewayImpl{
		ticketRegistryServiceURL: ticketRegistryServiceURL,
		getTicketsRoute:          getTicketsRoute,
		Logger:                   logger, // Pass logger to implementation
	}, nil
}

// GetTicketsForGivenTypeAndQuantity fetches tickets from the ticket registry service
func (c *TicketRegistryGatewayImpl) GetTicketsForGivenTypeAndQuantity(request *models.GetTicketsRequest) ([]models.Ticket, error) {
	// Log the start of fetching tickets
	c.Logger.WithFields(logrus.Fields{
		"ticketType": request.TicketType,
		"quantity":   request.Quantity,
	}).Info("Fetching tickets from ticket registry")

	// Build the full URL for the get-tickets API
	getTicketsURL := fmt.Sprintf("%s%s?ticketType=%s&quantity=%d", c.ticketRegistryServiceURL, c.getTicketsRoute, request.TicketType, request.Quantity)

	// Make the HTTP GET request
	resp, err := http.Get(getTicketsURL)
	if err != nil {
		c.Logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"url":   getTicketsURL,
		}).Error("Error making GET request to ticket registry")
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.Logger.WithFields(logrus.Fields{
				"error": err.Error(),
				"url":   getTicketsURL,
			}).Error("Error reading non-200 response body from ticket registry")
			return nil, err
		}
		c.Logger.WithFields(logrus.Fields{
			"statusCode": resp.StatusCode,
			"response":   string(body),
		}).Error("Ticket registry returned non-200 status")
		return nil, errors.New(string(body))
	}

	// Log success in fetching tickets
	c.Logger.WithFields(logrus.Fields{
		"url": getTicketsURL,
	}).Info("Successfully fetched tickets from ticket registry")

	// Parse the response body into []models.Ticket
	var tickets []models.Ticket
	if err := json.NewDecoder(resp.Body).Decode(&tickets); err != nil {
		c.Logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error decoding tickets response from ticket registry")
		return nil, err
	}

	// Return the tickets
	return tickets, nil
}

package gateways

import (
	"book-tickets/config"
	"book-tickets/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TicketRegistryGateway interface {
	GetTicketsForGivenTypeAndQuantity(request *models.GetTicketsRequest) ([]models.Ticket, error)
}

// TicketRegistryGatewayImpl is the concrete implementation of the TicketRegistryGateway interface
type TicketRegistryGatewayImpl struct {
	ticketRegistryServiceURL string
	getTicketsRoute          string
}

// NewTicketRegistryGateway creates a new instance of TicketRegistryGatewayImpl
func NewTicketRegistryGateway() (TicketRegistryGateway, error) {
	ticketRegistryServiceURL, err := config.ReadServiceConfig("ticketRegistry")
	if err != nil {
		return nil, err
	}

	getTicketsRoute, err := config.ReadAPIConfig("get-tickets")
	if err != nil {
		return nil, err
	}

	return &TicketRegistryGatewayImpl{
		ticketRegistryServiceURL: ticketRegistryServiceURL,
		getTicketsRoute:          getTicketsRoute,
	}, nil
}

// GetTicketsForGivenTypeAndQuantity fetches tickets from the ticketRegistry service
func (c *TicketRegistryGatewayImpl) GetTicketsForGivenTypeAndQuantity(request *models.GetTicketsRequest) ([]models.Ticket, error) {
	// Build the full URL for the get-tickets API
	getTicketsURL := fmt.Sprintf("%s%s?ticketType=%s&quantity=%d", c.ticketRegistryServiceURL, c.getTicketsRoute, request.TicketType, request.Quantity)

	// Make the HTTP GET request
	resp, err := http.Get(getTicketsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle non-200 responses
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	// Parse the response body into []models.Ticket
	var tickets []models.Ticket
	if err := json.NewDecoder(resp.Body).Decode(&tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

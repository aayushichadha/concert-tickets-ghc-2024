package service

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ticket-registry/mappers"
	"ticket-registry/models"
	"ticket-registry/repository"
)

func GetTicketsForGivenTypeAndQuantity(
	db *gorm.DB,
	getTicketsRequest *models.GetTicketsRequest,
	logger *logrus.Logger,
) (response *[]models.Ticket, err error) {
	logEntry := logrus.WithFields(logrus.Fields{
		"ticketType": getTicketsRequest.TicketType,
		"quantity":   getTicketsRequest.Quantity,
	})

	ticketsRepo := &repository.TicketRepository{DB: db} // Initialize with actual repository

	ticketTypeKey := mappers.AdaptToTicketTypeKey(getTicketsRequest.TicketType)

	tickets, err := ticketsRepo.GetTickets(string(ticketTypeKey), getTicketsRequest.Quantity)
	// If no ticket is found, return an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logEntry.WithField("error", err.Error()).Error("tickets not available")
			return nil, fmt.Errorf("ticket type '%s' with quantity %d not available", getTicketsRequest.TicketType, getTicketsRequest.Quantity)
		}
		return nil, err
	}

	// If the ticket exists and quantity is sufficient, deduct the requested quantity
	tickets.CurrentQuantity -= getTicketsRequest.Quantity

	updateErr := ticketsRepo.UpdateTickets(tickets)
	if updateErr != nil {
		logEntry.WithField("error", err.Error()).Error("failed to update ticket quantity")
		return nil, fmt.Errorf("failed to update ticket quantity: %v", updateErr)
	}

	// generate tickets as per quantity and type
	resp, _ := mappers.AdaptToTicketListFormat(getTicketsRequest)

	logEntry.Info("Processed get-tickets request")

	return resp, nil
}

func ShowTickets(db *gorm.DB, logger *logrus.Logger) (response *[]models.Tickets, err error) {
	ticketsRepo := &repository.TicketRepository{DB: db} // Initialize with actual repository

	tickets, err := ticketsRepo.GetAllTickets()
	// If no ticket is found, return an error
	if err != nil {
		logrus.WithField("error", err.Error()).Error("tickets not available")
		return nil, err
	}

	return tickets, nil
}

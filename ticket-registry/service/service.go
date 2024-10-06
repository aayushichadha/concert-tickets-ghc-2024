package service

import (
	"errors"
	"fmt"
	// "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	// "logging/logging"
	"ticket-registry/mappers"
	"ticket-registry/models"
	"ticket-registry/repository"
)

func GetTicketsForGivenTypeAndQuantity(db *gorm.DB, getTicketsRequest *models.GetTicketsRequest) (response *[]models.Ticket, err error) {
	ticketsRepo := &repository.TicketRepository{DB: db} // Initialize with actual repository

	ticketTypeKey := mappers.AdaptToTicketTypeKey(getTicketsRequest.TicketType)
	log.Printf("ticketTypeKey value %s", ticketTypeKey)

	tickets, err := ticketsRepo.GetTickets(string(ticketTypeKey), getTicketsRequest.Quantity)
	// If no ticket is found, return an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ticket type '%s' with quantity %d not available", getTicketsRequest.TicketType, getTicketsRequest.Quantity)
		}
		return nil, err
	}

	// If the ticket exists and quantity is sufficient, deduct the requested quantity
	tickets.CurrentQuantity -= getTicketsRequest.Quantity

	updateErr := ticketsRepo.UpdateTickets(tickets)
	if updateErr != nil {
		return nil, fmt.Errorf("failed to update ticket quantity: %v", updateErr)
	}

	// generate tickets as per quantity and type
	resp, _ := mappers.AdaptToTicketListFormat(getTicketsRequest)
	return resp, nil
}

func ShowTickets(db *gorm.DB) (response *[]models.Tickets, err error) {
	ticketsRepo := &repository.TicketRepository{DB: db} // Initialize with actual repository

	tickets, err := ticketsRepo.GetAllTickets()
	// If no ticket is found, return an error
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

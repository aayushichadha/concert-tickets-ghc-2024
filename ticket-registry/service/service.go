package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"ticket-registry/models"
	"ticket-registry/repository"
)

func GetTicketsForGivenTypeAndQuantity(db *gorm.DB, getTicketsRequest *models.GetTicketsRequest) (response *[]models.Ticket, err error) {
	ticketsRepo := &repository.TicketRepository{DB: db} // Initialize with actual repository

	tickets, err := ticketsRepo.GetTickets(getTicketsRequest.TicketType, getTicketsRequest.Quantity)
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
	resp, _ := adaptToTicketListFormat(getTicketsRequest)
	return resp, nil
}

func adaptToTicketListFormat(getTicketsRequest *models.GetTicketsRequest) (response *[]models.Ticket, err error) {
	var resp []models.Ticket
	for i := 0; i < getTicketsRequest.Quantity; i++ {
		ticket := models.Ticket{
			TicketID:   rand.Intn(100),
			TicketType: getTicketsRequest.TicketType,
		}
		resp = append(resp, ticket)
	}
	return &resp, nil
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

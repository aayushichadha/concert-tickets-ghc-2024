package service

import (
	"gorm.io/gorm"
	"math/rand"
	"ticket-registry/models"
	"ticket-registry/repository"
)

func GetTicketsForGivenTypeAndQuantity(db *gorm.DB, getTicketsRequest *models.GetTicketsRequest) (response *[]models.Ticket, err error) {
	ticketsRepo := &repository.TicketRepository{DB: db} // Initialize with actual repository

	_, err = ticketsRepo.GetTickets(getTicketsRequest.TicketType, getTicketsRequest.Quantity)
	if err != nil {
		return nil, err
	}

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

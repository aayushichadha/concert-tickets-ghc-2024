package service

import (
	"book-tickets/models"
	"gorm.io/gorm"
)

func BookTickets(db *gorm.DB, bookTicketsRequest *models.BookTicketsRequest) (response *[]models.Ticket, err error) {
	// add logic for booking tickets
	var resp []models.Ticket
	return &resp, nil
}

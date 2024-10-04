package repository

import (
	"gorm.io/gorm"
	"log"
	"ticket-registry/models"
)

type TicketRepository struct {
	DB *gorm.DB
}

// GetCatalogByProductID retrieves and catalog record by product ID from the database.
func (repo *TicketRepository) GetTickets(ticketType string, quantity int) (*models.Tickets, error) {
	var tickets models.Tickets
	// Query the tickets table based on the type and current quantity
	err := repo.DB.Where("ticket_type = ? AND current_quantity >= ?", ticketType, quantity).First(&tickets).Error

	log.Printf("GetTickets from database: tickets=%s, error=%s", tickets, err)
	return &tickets, err
}

func (repo *TicketRepository) UpdateTickets(tickets *models.Tickets) error {

	err := repo.DB.Save(tickets).Error
	return err
}

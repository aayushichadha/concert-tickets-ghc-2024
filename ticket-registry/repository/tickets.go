package repository

import (
	"gorm.io/gorm"
	"ticket-registry/models"
)

type TicketRepository struct {
	DB *gorm.DB
}

// GetTickets retrieves and catalog record by product ID from the database.
func (repo *TicketRepository) GetTickets(ticketType string, quantity int) (*models.Tickets, error) {
	var tickets models.Tickets
	// Query the tickets table based on the type and current quantity
	err := repo.DB.Where("ticket_type = ? AND current_quantity >= ?", ticketType, quantity).First(&tickets).Error
	return &tickets, err
}

func (repo *TicketRepository) UpdateTickets(tickets *models.Tickets) error {

	err := repo.DB.Save(tickets).Error
	return err
}

func (repo *TicketRepository) GetAllTickets() (*[]models.Tickets, error) {
	var tickets []models.Tickets

	// Fetch all records from the tickets table
	err := repo.DB.Find(&tickets).Error

	return &tickets, err
}

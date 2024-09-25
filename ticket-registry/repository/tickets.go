package repository

import (
	"ticket-registry/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

// CreateCatalog inserts a new catalog record into the database.
// func (repo *TicketRepository) CreateCatalog(catalog *models.Catalog) error {
// 	result := repo.DB.Create(catalog)
// 	return result.Error
// }

// // DeleteCatalog deletes an catalog record from the database.
// func (repo *TicketRepository) DeleteCatalog(catalog *models.Catalog) error {
// 	result := repo.DB.Delete(catalog)
// 	return result.Error
// }

// // UpdateCatalog updates an catalog record in the database.
// func (repo *TicketRepository) UpdateCatalog(catalog *models.Catalog) error {
// 	result := repo.DB.Save(catalog)
// 	return result.Error
// }

// // GetCatalog retrieves and catalog record by product ID from the database.
// func (repo *TicketRepository) GetCatalog() ([]models.Catalog, error) {
// 	var catalog []models.Catalog
// 	result := repo.DB.Find(&catalog)

// 	return catalog, result.Error
// }

// GetCatalogByProductID retrieves and catalog record by product ID from the database.
func (repo *TicketRepository) GetTickets(ticketType string, quantity int) (*models.Tickets, error) {
	var tickets models.Tickets
	 // Query the tickets table based on the type and current quantity
	err := repo.DB.Where("ticket_type = ? AND current_quantity >= ?", ticketType, quantity).Find(&tickets).Error

	return &tickets, err
}

package service

import (
	"book-tickets/gateways"
	"book-tickets/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

// BookingService is the service for handling ticket booking
type BookingService struct {
	CatalogGateway gateways.TicketRegistryGateway
	PaymentGateway gateways.PaymentGateway
}

// BookTickets implements the business logic for booking tickets
func (s *BookingService) BookTickets(db *gorm.DB, bookTicketsRequest *models.BookTicketsRequest) (response *[]models.TicketBooking, err error) {
	var bookedTickets []models.TicketBooking

	getTicketsRequest := &models.GetTicketsRequest{
		TicketType: bookTicketsRequest.Tickets.Type, // Assuming singular "Ticket"
		Quantity:   bookTicketsRequest.Tickets.Quantity,
	}

	tickets, err := s.CatalogGateway.GetTicketsForGivenTypeAndQuantity(getTicketsRequest)
	if err != nil {
		log.Println("Error fetching tickets:", err)
		return nil, err
	}

	// Checking if enough tickets are available
	if len(tickets) < bookTicketsRequest.Tickets.Quantity {
		return nil, errors.New("not enough tickets available")
	}

	paymentRequest := &models.MakePaymentRequest{
		UserID:        bookTicketsRequest.User.Id,
		Amount:        calculateTotalPrice(tickets),
		PaymentMethod: bookTicketsRequest.PaymentMethod,
	}

	err = s.PaymentGateway.MakePayment(paymentRequest)
	if err != nil {
		log.Println("Error processing payment:", err)
		return nil, err
	}

	tx := db.Begin()
	if tx.Error != nil {
		log.Println("Error starting transaction:", tx.Error)
		return nil, tx.Error
	}

	for _, ticket := range tickets {
		bookedTicket := models.TicketBooking{
			ID:     ticket.TicketID,
			UserID: bookTicketsRequest.User.Id,
			Type:   ticket.TicketType,
			Price:  ticket.Price,
		}

		// Save the booked ticket to the database
		if err := tx.Create(&bookedTicket).Error; err != nil {
			log.Println("Error saving booked ticket:", err)
			tx.Rollback()
			return nil, err
		}
		bookedTickets = append(bookedTickets, bookedTicket)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, err
	}

	// Return the booked tickets
	return &bookedTickets, nil
}

// Helper function to calculate the total price of the tickets
func calculateTotalPrice(tickets []models.Ticket) float64 {
	total := 0.0
	for _, ticket := range tickets {
		total += ticket.Price
	}
	return total
}

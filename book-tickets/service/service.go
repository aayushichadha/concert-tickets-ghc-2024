package service

import (
	"book-tickets/gateways"
	"book-tickets/models"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// BookingService is the service for handling ticket booking
type BookingService struct {
	CatalogGateway gateways.TicketRegistryGateway
	PaymentGateway gateways.PaymentGateway
	Logger         *logrus.Logger
}

// BookTickets implements the business logic for booking tickets
func (s *BookingService) BookTickets(ctx context.Context, db *gorm.DB, logger *logrus.Logger, bookTicketsRequest *models.BookTicketsRequest) (response *[]models.TicketBooking, err error) {
	var bookedTickets []models.TicketBooking

	// Create a reusable log entry with user details and ticket info
	logEntry := logger.WithFields(logrus.Fields{
		"user_id":    bookTicketsRequest.User.Id,
		"ticketType": bookTicketsRequest.Tickets.Type,
		"quantity":   bookTicketsRequest.Tickets.Quantity,
	})

	// Log start of ticket booking
	logEntry.Info("Starting ticket booking process")

	// Fetch available tickets
	getTicketsRequest := &models.GetTicketsRequest{
		TicketType: bookTicketsRequest.Tickets.Type,
		Quantity:   bookTicketsRequest.Tickets.Quantity,
	}

	tickets, err := s.CatalogGateway.GetTicketsForGivenTypeAndQuantity(ctx, getTicketsRequest)
	if err != nil {
		logEntry.Error("Error fetching tickets", err)
		return nil, err
	}

	// Check if enough tickets are available
	if len(tickets) < bookTicketsRequest.Tickets.Quantity {
		logEntry.WithFields(logrus.Fields{
			"requested": bookTicketsRequest.Tickets.Quantity,
			"available": len(tickets),
		}).Warn("Not enough tickets available")
		return nil, errors.New("not enough tickets available")
	}

	// Process payment
	paymentRequest := &models.MakePaymentRequest{
		UserID:        bookTicketsRequest.User.Id,
		Amount:        calculateTotalPrice(tickets),
		PaymentMethod: bookTicketsRequest.PaymentMethod,
	}

	err = s.PaymentGateway.MakePayment(paymentRequest)
	if err != nil {
		logEntry.WithFields(logrus.Fields{
			"amount":      paymentRequest.Amount,
			"paymentType": paymentRequest.PaymentMethod.Type,
		}).Error("Error processing payment", err)
		return nil, err
	}

	// Save booked tickets
	tx := db.Begin()
	for _, ticket := range tickets {
		bookedTicket := models.TicketBooking{
			ID:     ticket.TicketID,
			UserID: bookTicketsRequest.User.Id,
			Type:   ticket.TicketType,
			Price:  ticket.Price,
		}

		if err := tx.Create(&bookedTicket).Error; err != nil {
			logEntry.WithField("ticket_id", bookedTicket.ID).Error("Error saving booked ticket", err)
			tx.Rollback()
			return nil, err
		}
		bookedTickets = append(bookedTickets, bookedTicket)
	}

	if err := tx.Commit().Error; err != nil {
		logEntry.Error("Error committing transaction", err)
		return nil, err
	}

	// Log successful ticket booking
	logEntry.WithField("tickets_booked", len(bookedTickets)).Info("Tickets successfully booked")

	return &bookedTickets, nil
}

// Helper function to calculate total price
func calculateTotalPrice(tickets []models.Ticket) float64 {
	total := 0.0
	for _, ticket := range tickets {
		total += ticket.Price
	}
	return total
}

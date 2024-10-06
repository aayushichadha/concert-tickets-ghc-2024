package models

// User defines the details of the user making the booking
type User struct {
	Id   string `json:"id"`   // User's unique identifier
	Name string `json:"name"` // User's name
	DOB  string `json:"dob"`  // User's date of birth (format: YYYY-MM-DD)
}

// BookTicketsRequest defines the complete booking request
type BookTicketsRequest struct {
	User          User          `json:"user"`           // User details
	Tickets       Tickets       `json:"ticket"`         // Ticket details
	PaymentMethod PaymentMethod `json:"payment_method"` // Payment method details
}

type Tickets struct {
	Type     string `json:"type"`     // Type of ticket (e.g., VIP, Regular)
	Quantity int    `json:"quantity"` // Number of tickets to book
}

// TicketBooking represents a ticket that has been booked by a user
type TicketBooking struct {
	ID     string  `json:"id"`      // Ticket ID (from catalog)
	UserID string  `json:"user_id"` // User who booked the ticket
	Type   string  `json:"type"`    // Type of ticket (e.g., VIP, Regular)
	Price  float64 `json:"price"`   // Price of the ticket
}

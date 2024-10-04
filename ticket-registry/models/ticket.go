package models

type GetTicketsRequest struct {
	TicketType string
	Quantity   int
}

type Ticket struct {
	TicketID   int
	TicketType string
}

type Tickets struct {
	ID              int
	TicketType      string
	CurrentQuantity int
}

package models

type GetTicketsRequest struct {
	TicketType string
	Quantity   int
}

type Ticket struct {
	TicketID   string
	TicketType string
	Price      float64
}

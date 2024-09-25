package models

type GetTicketsRequest struct {
	TicketType        string
	Quantity   int
}

type Ticket struct {
	TicketID   int
	TicketType        string
}

type Tickets struct {
	TicketType        string
	CurrentQuantity        int
}

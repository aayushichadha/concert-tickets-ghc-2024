package models

type GetTicketsRequest struct {
	TicketType string
	Quantity   int
}

type Ticket struct {
	TicketID   string
	TicketType string
}

type Tickets struct {
	ID              int
	TicketType      TicketType
	CurrentQuantity int
}

// Define TicketType as its own type
type TicketType string

// Define possible values for TicketType
const (
	VIPFrontRow       TicketType = "VIPFrontRow"
	PlatinumSeats     TicketType = "PlatinumSeats"
	GeneralAdmissions TicketType = "GeneralAdmissions"
	BalconySeat       TicketType = "BalconySeat"
	SuperfanPit       TicketType = "SuperfanPit"
)

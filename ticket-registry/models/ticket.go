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
	TicketType      TicketType
	CurrentQuantity int
}

// Define TicketType as its own type
type TicketType string

// Define possible values for TicketType
const (
	VIPFrontRow      TicketType = "VIPFrontRow"
	PlatinumSeating  TicketType = "PlatinumSeating"
	GeneralAdmission TicketType = "GeneralAdmission"
	BalconySeating   TicketType = "BalconySeating"
	SuperfanPit      TicketType = "SuperfanPit"
)

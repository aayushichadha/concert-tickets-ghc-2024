package models

type BookTicketsRequest struct {
	User struct {
		Id   string `json:"Id"`
		Name string `json:"Name"`
		DOB  string `json:"DOB"`
	} `json:"User"`
	Ticket struct {
		Type     string
		Quantity int
	} `json:"Ticket"`
	PaymentMethod struct {
		Type          string `json:"Type"`
		Number        string `json:"Number"`
		Authorization string `json:"Authorization"`
	} `json:"PaymentMethod"`
}

type Ticket struct {
	TicketID   int
	TicketType string
}

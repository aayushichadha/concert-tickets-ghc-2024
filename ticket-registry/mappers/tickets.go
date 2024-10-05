package mappers

import (
	"math/rand"
	"strconv"
	"ticket-registry/models"
)

func AdaptToTicketListFormat(getTicketsRequest *models.GetTicketsRequest) (response *[]models.Ticket, err error) {
	var resp []models.Ticket
	for i := 0; i < getTicketsRequest.Quantity; i++ {
		ticket := models.Ticket{
			TicketID:   strconv.Itoa(rand.Intn(100)),
			TicketType: getTicketsRequest.TicketType,
		}
		resp = append(resp, ticket)
	}
	return &resp, nil
}

func AdaptToTicketTypeKey(ticketType string) models.TicketType {
	switch ticketType {
	case "vip-front-row":
		return models.VIPFrontRow
	case "platinum-seating":
		return models.PlatinumSeating
	case "general-admission":
		return models.GeneralAdmission
	case "balcony-seating":
		return models.BalconySeating
	case "superfan-pit":
		return models.SuperfanPit
	default:
		return ""
	}
}

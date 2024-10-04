package models

type MakePaymentRequest struct {
	UserID        string        `json:"user_id"` // ID of the user making the payment
	Amount        float64       `json:"amount"`
	PaymentMethod PaymentMethod `json:"PaymentMethod"`
}

type PaymentMethod struct {
	Type          string `json:"Type"`
	Number        string `json:"Number"`
	Authorization string `json:"Authorization"`
}

type MakePaymentResponse struct {
	PaymentId string `json:"PaymentId"`
	Status    string `json:"Status"`
}

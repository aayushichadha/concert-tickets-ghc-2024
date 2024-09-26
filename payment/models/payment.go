package models

type MakePaymentRequest struct {
	User struct {
		Id   string `json:"Id"`
		Name string `json:"Name"`
		DOB  string `json:"DOB"`
	} `json:"User"`
	PaymentMethod struct {
		Type          string `json:"Type"`
		Number        string `json:"Number"`
		Authorization string `json:"Authorization"`
	} `json:"PaymentMethod"`
}

type MakePaymentResponse struct {
	PaymentId string `json:"PaymentId"`
	Status    string `json:"Status"`
}

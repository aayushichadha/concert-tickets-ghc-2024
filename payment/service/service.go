package service

import (
	"log"
	"payment/models"
)

func InitiatePayment(payment models.Payment) (err error) {
	return nil
}

func RollbackPayment(payment models.Payment) (err error) {
	log.Printf("Rolling back payment with ID %d", payment.ID)

	return nil
}

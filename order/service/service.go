package service

import (
	"gorm.io/gorm"
	"order/models"
)

func PlaceOrder(db *gorm.DB, customerID, productID, quantity int) (order models.Order, err error) {
	return order, nil
}

func updateCatalog(productID int, quantity int) (err error) {
	return nil
}

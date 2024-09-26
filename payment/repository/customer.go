package repository

import (
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

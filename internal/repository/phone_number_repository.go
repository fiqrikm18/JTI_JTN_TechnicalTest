package repository

import (
	"gorm.io/gorm"
)

type PhoneNumberRepository struct {
	DB *gorm.DB
}

func NewPhoneNumberRepository(db *gorm.DB) *PhoneNumberRepository {
	return &PhoneNumberRepository{
		DB: db,
	}
}

func (repo *PhoneNumberRepository) Insert(phoneNumber interface{}) error {
	return repo.DB.Create(phoneNumber).Error
}

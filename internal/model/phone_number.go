package model

import (
	"gorm.io/gorm"
	"time"
)

type PhoneNumber struct {
	ID          uint   `json:"id" gorm:"primaryKey" sql:"AUTO_INCREMENT"`
	PhoneNumber string `json:"phone_number"`
	Provider    string `json:"provider"`
	Type        int    `json:"type"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type PhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
	Provider    string `json:"provider"`
}

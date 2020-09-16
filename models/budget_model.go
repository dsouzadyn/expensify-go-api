package models

import (
	"time"

	"gorm.io/gorm"
)

// Budget model
type Budget struct {
	gorm.Model
	Name           string    `json:"name" gorm:"not null"`
	Description    string    `json:"description" gorm:"not null"`
	Amount         int64     `json:"amount" gorm:"not null"`
	StartDate      time.Time `json:"startdate" gorm:"not null"`
	EndDate        time.Time `json:"enddate" gorm:"not null"`
	UserID         int
	User           User
	ExchangeRateID int
	ExchangeRate   ExchangeRate
}

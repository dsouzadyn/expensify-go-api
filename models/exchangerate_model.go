package models

import "gorm.io/gorm"

// ExchangeRate model
type ExchangeRate struct {
	gorm.Model
	Currency string `json:"currency" gorm:"not null"`
	Rate     int64  `json:"rate" gorm:"not null"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

// Expense model
type Expense struct {
	gorm.Model
	Name           string    `json:"name" gorm:"not null"`
	Description    string    `json:"description" gorm:"not null"`
	Amount         int64     `json:"amount" gorm:"not null"`
	Date           time.Time `json:"date" gorm:"not null"`
	BudgetID       int
	Budget         Budget
	ExchangeRateID int
	ExchangeRate   ExchangeRate
	CategoryID     int
	Category       Category
}

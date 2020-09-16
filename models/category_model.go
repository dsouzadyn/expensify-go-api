package models

import "gorm.io/gorm"

// Category model
type Category struct {
	gorm.Model
	Name string `json:"category" gorm:"not null"`
}

package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model declared below
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

// BeforeCreate User hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	passwordHash, err := HashPassword(u.Password)
	if err != nil {
		return errors.New("could not create password hash")
	}
	u.Password = passwordHash
	return nil
}

// HashPassword returns a hash for the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares the password and hash to verify
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

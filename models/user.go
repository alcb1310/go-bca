package models

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
	"gitlab.com/0x4149/logz"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name" gorm:"index"`

	CompanyId uuid.UUID `json:"companyId" gorm:"not null;type:uuid"`
	Company   Company   `json:"company" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(l LoginType, db *gorm.DB) (*User, error) {
	if _, err := mail.ParseAddress(l.Email); err != nil {
		logz.Debug("Invalid email address")
		return nil, errors.New("Invalid credentials")
	}
	var user User
	result := db.Find(&user, "email = ?", l.Email)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected != 1 {
		return nil, errors.New("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password)); err != nil {
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:now()"`
}

type Company struct {
	Model
	Ruc       string `json:"ruc" gorm:"unique;not null"`
	Name      string `json:"name" gorm:"unique;not null"`
	Employees uint   `json:"employees" gorm:"not null;type:int;default:1"`
	IsActive  bool   `json:"isActive" gorm:"not null;default:true"`
}

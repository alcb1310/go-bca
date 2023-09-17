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

type User struct {
	Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Name     string `json:"name" gorm:"index"`

	CompanyId string  `json:"companyId" gorm:"not null;type:uuid"`
	Company   Company `json:"company" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

type Project struct {
	Model
	Name     string `json:"name" gorm:"not null;index;index:uq_projectname,unique"`
	IsActive bool   `json:"isActive" gorm:"not null;default:true"`

	Company   Company `json:"company" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CompanyId string  `json:"companyId" gorm:"not null;type:uuid;index:uq_projectname,unique"`

	User   User   `json:"user" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	UserId string `json:"userId" gorm:"not null;type:uuid"`
}

type Supplier struct {
	Model
	SupplierId   string `json:"supplier_id" gorm:"not null;index:uq_supplierid,unique"`
	Name         string `json:"name" gorm:"not null;index:uq_suppliername,unique;index"`
	ContactName  string `json:"contact_name"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`

	Company   Company `json:"company" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CompanyId string  `json:"companyId" gorm:"not null;type:uuid;index:uq_supplierid,unique;index:uq_suppliername,unique"`

	User   User   `json:"user" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	UserId string `json:"userId" gorm:"not null;type:uuid"`
}

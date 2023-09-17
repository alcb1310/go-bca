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

type BudgetItem struct {
	Model
	Code        string `json:"code" gorm:"not null;index;index:uq_budgetitemcode,unique"`
	Name        string `json:"name" gorm:"not null;index;index:uq_budgetitemname,unique"`
	Accumulates bool   `json:"accumulates" gorm:"not null;default:false"`
	Level       uint   `json:"level" gorm:"not null;type:uint;default:1"`

	BudgetItem   *BudgetItem `json:"budgetItem" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	BudgetItemId string      `json:"budgetItemId" gorm:"type:uuid"`

	Company   Company `json:"company" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CompanyId string  `json:"companyId" gorm:"not null;type:uuid;index:uq_budgetitemcode,unique;index:uq_budgetitemname,unique"`

	User   User   `json:"user" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	UserId string `json:"userId" gorm:"not null;type:uuid"`
}

type Budget struct {
	Model

	BudgetItem   BudgetItem `json:"budgetItem" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	BudgetItemId string     `json:"budgetItemId" gorm:"not null;type:uuid;index:uq_budgetprojectbudgetitem,unique"`
	Project      Project    `json:"project" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	ProjectId    string     `json:"projectId" gorm:"not null;type:uuid;index:uq_budgetprojectbudgetitem,unique"`

	InitialQuantity float64 `json:"initialQuantity" gorm:"type:decimal(20,8)"`
	InitialCost     float64 `json:"initialCost" gorm:"type:decimal(20,8)"`
	InitialTotal    float64 `json:"initialTotal" gorm:"not null;type:decimal(20,8)"`

	SpentQuantity float64 `json:"spentQuantity" gorm:"type:decimal(20,8)"`
	SpentTotal    float64 `json:"spentTotal" gorm:"not null;type:decimal(20,8)"`

	ToSpendQuantity float64 `json:"toSpendQuantity" gorm:"type:decimal(20,8)"`
	ToSpendCost     float64 `json:"toSpendCost" gorm:"type:decimal(20,8)"`
	ToSpendtTotal   float64 `json:"toSpendtTotal" gorm:"not null;type:decimal(20,8)"`

	UpdatedBudget float64 `json:"updated_budget" gorm:"not null;type:decimal(20,8)"`

	Company   Company `json:"company" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	CompanyId string  `json:"companyId" gorm:"not null;type:uuid;index:uq_budgetprojectbudgetitem,unique"`

	User   User   `json:"user" gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	UserId string `json:"userId" gorm:"not null;type:uuid"`
}

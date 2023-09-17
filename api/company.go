package api

import (
	"encoding/json"
	"net/http"

	"github.com/alcb1310/bca-go/models"
	"gorm.io/gorm"
)

type registerCompany struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Employees    uint   `json:"employees"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserFullName string `json:"user_full_name"`
}

var database *gorm.DB

func register(w http.ResponseWriter, r *http.Request) {
	var newCompany registerCompany
	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := models.Company{
		Ruc:       newCompany.Id,
		Name:      newCompany.Name,
		Employees: newCompany.Employees,
	}

	tx := database.Begin()

	if err := tx.Create(&c).Error; err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		tx.Rollback()
		return
	}

	u := models.User{
		Email:    newCompany.UserName,
		Name:     newCompany.UserFullName,
		Password: newCompany.UserPassword,
		Company:  c,
	}

	if err = tx.Create(&u).Error; err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		tx.Rollback()
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (b *Router) companyRoutes() {
	database = b.DB.Data
	b.r.HandleFunc("/register", register).Methods("POST")
}

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"time"

	"github.com/alcb1310/bca-go/constants"
	"github.com/alcb1310/bca-go/models"
	"github.com/alcb1310/bca-go/utils"
	"gitlab.com/0x4149/logz"
	"golang.org/x/crypto/bcrypt"
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

var secretKey = os.Getenv("SECRET")

// const secretKey = "MySuperSecretKey"

func register(w http.ResponseWriter, r *http.Request) {
	var newCompany registerCompany
	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := mail.ParseAddress(newCompany.UserName); err != nil {
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

	pass, _ := bcrypt.GenerateFromPassword([]byte(newCompany.UserPassword), constants.SALT)

	u := models.User{
		Email:    newCompany.UserName,
		Name:     newCompany.UserFullName,
		Password: string(pass),
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

func login(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginType
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "No email/password", http.StatusBadRequest)
		return
	}

	u, err := models.Login(credentials, database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	logz.Debug(secretKey)
	jwtMaker, err := utils.NewJWTMaker(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := jwtMaker.CreateToken(*u, 60*time.Minute)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(response{
		Message: fmt.Sprintf("Bearer %s", token),
	})
}

func (b *Router) companyRoutes() {
	database = b.DB.Data
	b.r.HandleFunc("/register", register).Methods(http.MethodPost)
	b.r.HandleFunc("/login", login).Methods(http.MethodPost)
}

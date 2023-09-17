package api

import (
	"encoding/json"
	"net/http"
)

type registerCompany struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Employees    uint   `json:"employees"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserFullName string `json:"user_full_name"`
}

func register(w http.ResponseWriter, r *http.Request) {
	var newCompany registerCompany
	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode("success")
}

func (b *Router) companyRoutes() {
	b.r.HandleFunc("/register", register).Methods("POST")
}

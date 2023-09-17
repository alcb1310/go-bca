package api

import (
	"encoding/json"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := response{
		Message: "This is the register route",
	}
	json.NewEncoder(w).Encode(res)
}

func (b *Router) companyRoutes() {
	b.r.HandleFunc("/register", register)

}

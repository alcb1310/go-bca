package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gitlab.com/0x4149/logz"
)

type errorResponse struct {
	Error string `json:"error"`
}

func authVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errorResponse{
				Error: "Missing autherization token",
			})
			return
		}

		token := strings.Split(bearerToken, " ")
		if len(token) != 2 {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errorResponse{
				Error: "Invalid autherization token",
			})
			return
		}
		logz.Debug(token[1])

		next.ServeHTTP(w, r)
	})
}

func logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(response{
		Message: "Log out",
	})
}

func authRoutes(r *mux.Router) {
	r.HandleFunc("/logout", logout).Methods(http.MethodGet)

}

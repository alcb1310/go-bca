package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/alcb1310/bca-go/models"
	"github.com/alcb1310/bca-go/utils"
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

		secretKey := os.Getenv("SECRET")
		maker, err := utils.NewJWTMaker(secretKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errorResponse{
				Error: "Unable to authenticate",
			})
			return
		}

		tokenData, err := maker.VerifyToken(token[1])
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errorResponse{
				Error: "Unable to authenticate",
			})
			return
		}

		marshalStr, _ := json.Marshal(tokenData)
		ctx := r.Context()
		ctx = context.WithValue(r.Context(), "token", marshalStr)
		r = r.Clone(ctx)
		var u models.LoggedInUser
		result := database.First(&u, "email=?", tokenData.Email)
		if result.Error != nil || result.RowsAffected != 1 {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errorResponse{
				Error: "Invalid token",
			})
			return
		}

		if !bytes.Equal([]byte(token[1]), u.JWT) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&errorResponse{
				Error: "Invalid token",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func logout(w http.ResponseWriter, r *http.Request) {
	token, err := GetMyPaload(r)
	if err != nil {
		logz.Error("Process fucked up")
		return
	}

	database.Delete(&models.LoggedInUser{}, "email = ?", token.Email)

	json.NewEncoder(w).Encode(response{
		Message: "Log out",
	})
}

func authRoutes(r *mux.Router) {
	r.HandleFunc("/logout", logout).Methods(http.MethodGet)

}

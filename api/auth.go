package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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
		logz.Error(err)
		return
	}

	database.Delete(&models.LoggedInUser{}, "email = ?", token.Email)

	json.NewEncoder(w).Encode(response{
		Message: "User logged out",
	})
}

func refresh(w http.ResponseWriter, r *http.Request) {
	oldToken, err := GetMyPaload(r)
	if err != nil {
		logz.Error(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var u models.User
	result := database.Find(&u, "email = ?", oldToken.Email)
	if result.Error != nil {
		logz.Error("This should never occur")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	secretKey := os.Getenv("SECRET")
	jwtMaker, err := utils.NewJWTMaker(secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := jwtMaker.CreateToken(u, 60*time.Minute)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	byteToken := []byte(token)
	loggedInUser := models.LoggedInUser{
		Email: u.Email,
		JWT:   byteToken,
	}

	database.Save(&loggedInUser)

	json.NewEncoder(w).Encode(response{
		Message: fmt.Sprintf("Bearer %s", token),
	})
}

func authRoutes(r *mux.Router) {
	r.HandleFunc("/logout", logout).Methods(http.MethodGet)
	r.HandleFunc("/refresh", refresh).Methods(http.MethodPut)

	// project routes
	r.HandleFunc("/projects", createProject).Methods(http.MethodPost)
	r.HandleFunc("/projects", getAllProjects).Methods(http.MethodGet)
}

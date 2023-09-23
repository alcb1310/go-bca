package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type contextPayload struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	CompanyId  uuid.UUID `json:"company_id"`
	IsLoggedIn bool      `json:"is_logged_in"`
	IssuedAt   time.Time `json:"issued_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}

func SetMyPayload(r *http.Request, p contextPayload) {}

func GetMyPaload(r *http.Request) (contextPayload, error) {
	ctx := r.Context()
	val := ctx.Value("token")

	x, ok := val.([]byte)
	if !ok {
		return contextPayload{}, errors.New("Unable to load context")
	}
	var p contextPayload
	err := json.Unmarshal(x, &p)
	if err != nil {
		return contextPayload{}, errors.New("Unable to parse context")
	}
	return p, nil
}

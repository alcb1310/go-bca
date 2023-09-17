package api

import (
	"encoding/json"
	"net/http"

	"github.com/alcb1310/bca-go/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type response struct {
	Message string `json:"message"`
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type Router struct {
	DB      models.DB
	r       *mux.Router
	Handler http.Handler
}

func (b *Router) Routes() {
	b.DB.Initialize()
	b.r = mux.NewRouter()

	b.r.Use(jsonMiddleware) // All responses will be of type application/json

	b.r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		res := response{
			Message: "Last test",
		}
		json.NewEncoder(w).Encode(res)
	})
	b.companyRoutes()
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	b.Handler = handlers.CORS(originsOk, methodsOk)(b.r)
}

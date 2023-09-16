package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gitlab.com/0x4149/logz"
)

type response struct {
	Message string `json:"message"`
}

func init() {
	logz.VerbosMode()

	logz.Run()
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	port, portRead := os.LookupEnv("PORT")
	if !portRead {
		godotenv.Load()
		port, portRead = os.LookupEnv("PORT")
		if !portRead {
			logz.Fatal("Unable to load environment variables")
		}
	}
	r := mux.NewRouter()
	r.Use(jsonMiddleware) // All responses will be of type application/json

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		res := response{
			Message: "Hello World!!!!",
		}
		json.NewEncoder(w).Encode(res)
	})

	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	hand := handlers.CORS(originsOk, methodsOk)(r)

	logz.Info("Server Running...\n")
	logz.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), hand))
}

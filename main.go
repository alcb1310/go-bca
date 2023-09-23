package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alcb1310/bca-go/api"
	"github.com/joho/godotenv"
	"gitlab.com/0x4149/logz"
)

// Initialize the logging system with logz
func init() {
	logz.VerbosMode()
	logz.Run()
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
	r := api.Router{}
	r.Routes()

	logz.Info("Server Running...\n")
	logz.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r.Handler))
}

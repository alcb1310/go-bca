package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/alcb1310/bca-go/api"
	"github.com/joho/godotenv"
)

var a api.Router

func TestMain(m *testing.M) {
	fmt.Println("Test starting")
	cmd := exec.Command("/bin/bash", "-c", "./resetDB.sh")
	status, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(status))
	fmt.Println("Database created")

	godotenv.Load(".env.test")

	r := api.Router{}
	r.Routes()

	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Routes()

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestEmptyTable(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:8080/invalid", nil)
	if err != nil {
		fmt.Println("Error with the request")
		return
	}
	// fmt.Println(req.URL)
	response := executeRequest(req)
	fmt.Println(response)

	// checkResponseCode(t, http.StatusNotFound, response.Code)
}

package api

import (
	"encoding/json"
	"net/http"

	"github.com/alcb1310/bca-go/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gitlab.com/0x4149/logz"
)

type projectEdit struct {
	Name string `json:"name"`
}
type project struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CompanyId uuid.UUID `json:"company_id"`
}

func createProject(w http.ResponseWriter, r *http.Request) {
	var pro projectEdit
	var user models.User
	var company models.Company

	payload, err := GetMyPaload(r)
	if err != nil {
		logz.Error(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&pro)

	database.Find(&user, "email = ?", payload.Email)
	database.Find(&company, "id = ?", payload.CompanyId)
	createdProject := models.Project{
		Name:    pro.Name,
		User:    user,
		Company: company,
	}

	result := database.Create(&createdProject)
	if result.Error != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response{
			Message: "Project already exists",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProject)
}

func getAllProjects(w http.ResponseWriter, r *http.Request) {
	var allProjects []project

	payload, err := GetMyPaload(r)
	if err != nil {
		logz.Error(err)
		return
	}
	database.Find(&allProjects, "company_id = ?", payload.CompanyId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allProjects)
}

func getOneProject(w http.ResponseWriter, r *http.Request) {
	var selectedProject project
	vars := mux.Vars(r)
	projectId := vars["projectId"]

	payload, err := GetMyPaload(r)
	if err != nil {
		logz.Error(err)
		return
	}

	result := database.Find(&selectedProject, "company_id = ? and id = ?", payload.CompanyId, projectId)
	if result.Error != nil || result.RowsAffected != 1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{
			Message: "Project not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(selectedProject)
}

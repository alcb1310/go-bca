package api

import (
	"encoding/json"
	"net/http"

	"github.com/alcb1310/bca-go/models"
	"github.com/google/uuid"
	"gitlab.com/0x4149/logz"
)

type projectEdit struct {
	Name string `json:"name"`
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
		// http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProject)
}

func getAllProjects(w http.ResponseWriter, r *http.Request) {
	type Project struct {
		Id        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		IsActive  bool      `json:"is_active"`
		CompanyId uuid.UUID `json:"company_id"`
	}
	var allProjects []Project

	payload, err := GetMyPaload(r)
	if err != nil {
		logz.Error(err)
		return
	}
	logz.Debug(payload.CompanyId)
	database.Find(&allProjects, "company_id = ?", payload.CompanyId)
	// database.Model(&models.Project{}).Select("\"project\".name, \"project\".is_active").Joins("left join company on company.id = projects.company_id").Where("company.id = ?", payload.CompanyId).Find(&allProjects)
	// database.Joins("Company").Select("\"project\".name, \"project\".is_active").Find(&allProjects)
	// database.Model(models.Project{CompanyId: payload.CompanyId}).Find(&allProjects)
	// database.Raw("select name, is_active from project where company_id = ?", payload.CompanyId).Scan(&allProjects)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allProjects)
}

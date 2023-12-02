package rest

import (
	"backendAleshinaFeklistova/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreateProject
//
//	@Summary	New Project Creating
//	@ID			createProject
//	@Tags		Project
//	@Accept		mpfd
//	@Param		Dataset			formData	file	true	"Dataset file"
//	@Param		ProjectName		formData	string	true	"project name"
//	@Param		ProjectType		formData	string	true	"project type"	Enums(Логистическая регрессия, Случайный лес)
//	@Param		ProjectTarget	formData	string	true	"project info"
//	@Security		ApiKeyAuth
//	@Success	200				{string}	string	"ok"
//	@Router		/project/create [post]
func CreateProject(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	err := service.CreateProject(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occured while parsing form: %s", err)
	}
}

// GetProjectById
//
//	@Summary	Get project by id
//	@ID			getProjectById
//	@Tags		Project
//	@Param		id	path		int				true	"Project ID"
//	@Security		ApiKeyAuth
//	@Success	200	{object}	models.Project	"ok"
//	@Router		/project/{id} [get]
func GetProjectById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	res, err := json.Marshal(service.GetProjectById(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// GetAllProjects
//
//	@Summary	Get all projects
//	@ID			getAllProjects
//	@Tags		Project
//	@Security		ApiKeyAuth
//	@Success	200	{object}	[]models.Project	"ok"
//	@Router		/project [get]
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(service.GetAllProjects())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// EditProject
//
//	@Summary	Edit Project
//	@ID			editProject
//	@Tags		Project
//	@Param		Dataset			formData	file	true	"Dataset file"
//	@Param		ProjectName		formData	string	true	"project name"
//	@Param		ProjectType		formData	string	true	"project type"	Enums(Логистическая регрессия, Случайный лес, Линейная регрессия, Метод опорных векторов)
//	@Param		ProjectTarget	formData	string	true	"project info"
//	@Security		ApiKeyAuth
//	@Success	200	{string}	string	"ok"
//	@Router		/project/edit [put]
func EditProject(w http.ResponseWriter, r *http.Request) {
	err := service.EditProject(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// DeleteProject
//
//	@Summary	Delete Project
//	@ID			deleteProject
//	@Tags		Project
//	@Param		id	path		int		true	"Project ID"
//	@Security		ApiKeyAuth
//	@Success	200	{string}	string	"ok"
//	@Router		/project/{id} [delete]
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	service.DeleteProject(id)
}

// GetProjectsByType
//
//	@Summary	Get Projects by type
//	@ID			getProjByType
//	@Tags		Project
//	@Param		type	path		string		true	"Project type" Enums(logreg, randf, linreg, vect)
//	@Security		ApiKeyAuth
//	@Success	200	{object}	[]models.Project	"ok"
//	@Router		/project/type/{type} [get]
func GetProjectsByType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	t, _ := params["type"]
	var prType string
	switch t {
	case "logreg":
		prType = "Логистическая регрессия"
	case "randf":
		prType = "Случайный лес"
	case "linreg":
		prType = "Линейная регрессия"
	case "vect":
		prType = "Метод опорных векторов"
	}
	res, err := json.Marshal(service.GetProjectsByType(prType))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

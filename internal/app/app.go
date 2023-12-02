package app

import (
	"backendAleshinaFeklistova/internal/service"
	"backendAleshinaFeklistova/internal/transport/rest"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
)

func (a *App) StartApp() error {
	log.Println("App starting")
	a.initRouter()
	log.Printf("Server started on :%s", a.port)
	err := http.ListenAndServe(a.host+":"+a.port, a.r)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initRouter() {
	a.r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://"+a.host+":"+a.port+"/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	a.r.HandleFunc("/user/register", rest.RegisterUser).Methods(http.MethodPost)
	a.r.HandleFunc("/user/login", rest.Login).Methods(http.MethodPost)
	a.r.HandleFunc("/user/{id}", service.VerifyJWT(rest.GetUserById)).Methods(http.MethodGet)
	a.r.HandleFunc("/user/edit", service.VerifyJWT(rest.EditProfile)).Methods(http.MethodPut)
	a.r.HandleFunc("/user/{id}", service.VerifyJWT(rest.DeleteProfile)).Methods(http.MethodDelete)
	a.r.HandleFunc("/user", service.VerifyJWT(rest.GetAll)).Methods(http.MethodGet)

	a.r.HandleFunc("/project/create", service.VerifyJWT(rest.CreateProject)).Methods(http.MethodPost)
	a.r.HandleFunc("/project/{id}", service.VerifyJWT(rest.GetProjectById)).Methods(http.MethodGet)
	a.r.HandleFunc("/project", service.VerifyJWT(rest.GetAllProjects)).Methods(http.MethodGet)
	a.r.HandleFunc("/project/edit", service.VerifyJWT(rest.EditProject)).Methods(http.MethodPut)
	a.r.HandleFunc("/project/{id}", service.VerifyJWT(rest.DeleteProject)).Methods(http.MethodDelete)
	a.r.HandleFunc("/project/type/{type}", service.VerifyJWT(rest.GetProjectsByType)).Methods(http.MethodGet)

}

package app

import "github.com/gorilla/mux"

type Starter interface {
	StartApp() error
	initRouter()
}

type App struct {
	host string
	port string
	r    *mux.Router
}

func NewApp(host string, port string) *App {
	return &App{
		host: host,
		port: port,
		r:    mux.NewRouter(),
	}
}

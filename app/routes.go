package app

import (
	"net/http"

	"github.com/nandarahmat/app/controllers"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

	staticFIleDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFIleDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
package handlers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HealthCheckHandler).Methods("GET")

	router.HandleFunc("/inputInfo",reciveInfo).Methods("POST")


	return router
}

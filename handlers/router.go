package handlers

import (
	"github.com/gorilla/mux"

)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/", test).Methods("POST")
	router.HandleFunc("/getInfo",giveInfo).Methods("GET")
	router.HandleFunc("/inputInfo",reciveInfo).Methods("POST")

	router.Use(mux.CORSMethodMiddleware(router))
	return router
}

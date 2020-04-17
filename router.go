package main

import (
	"github.com/Turing00/go_cars_api/controller"
	"github.com/gorilla/mux"
)

// InitializeRouter function creates a new router object
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/cars").HandlerFunc(controller.FindAllCarsEndpoint)
	router.Methods("GET").Path("/cars/{id}").HandlerFunc(controller.FindCarEndpoint)
	router.Methods("POST").Path("/cars").HandlerFunc(controller.CreateCarEndpoint)
	router.Methods("PUT").Path("/cars/{id}").HandlerFunc(controller.UpdateCarEndpoint)
	router.Methods("DELETE").Path("/cars/{id}").HandlerFunc(controller.DeleteCarEndpoint)
	return router
}

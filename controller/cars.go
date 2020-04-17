package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Turing00/go_cars_api/dao"
	"github.com/Turing00/go_cars_api/model"
	"github.com/gorilla/mux"
)

// FindCarEndpoint function gets car from the table cars inside postgresDb instance named go_cars_api based on ID from json object
func FindCarEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	errorCheck(err)

	car := dao.FindByID(id)

	json.NewEncoder(w).Encode(car)
}

// FindAllCarsEndpoint function gets all cars from the table cars inside postgresDb instance named go_cars_api into json object
func FindAllCarsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	cars := dao.FindAll()

	json.NewEncoder(w).Encode(cars)
}

// CreateCarEndpoint function puts car into the table cars inside postgresDb instance named go_cars_api and from json object
func CreateCarEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	errorCheck(err)

	var (
		car model.Car
	)

	err = json.Unmarshal(body, &car)

	errorCheck(err)

	dao.Insert(&car)

	json.NewEncoder(w).Encode(car)
}

// UpdateCarEndpoint function updates car from the table cars inside postgresDb instance named go_cars_api and from json object
func UpdateCarEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	errorCheck(err)

	body, err := ioutil.ReadAll(r.Body)

	errorCheck(err)

	car := dao.FindByID(id)

	err = json.Unmarshal(body, &car)

	errorCheck(err)

	dao.Update(car)

	json.NewEncoder(w).Encode(car)
}

// DeleteCarEndpoint function deletes car from the table cars inside postgresDb instance named go_cars_api and from json object
func DeleteCarEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	errorCheck(err)

	err = dao.DeleteByID(id)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

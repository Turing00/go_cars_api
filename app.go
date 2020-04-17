package main

import (
	"log"
	"net/http"

	"github.com/Turing00/go_cars_api/dao"
)

func main() {

	dao.ConnectAndCreateTable()

	router := InitializeRouter()

	// Populate database
	//dao.Insert(&model.Car{Manufacturer: "citroen", Design: "ds3", Style: "sport", Doors: 4})

	err := http.ListenAndServe(":8080", router)

	log.Fatal(err)
}

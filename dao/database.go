package dao

import (
	"database/sql"
	"log"
	"time"

	"github.com/Turing00/go_cars_api/model"
	// Avoid any side effect from lib/pq library
	_ "github.com/lib/pq"
)

// carsDAO struct deals with postgres Db instance credentials and more
type carsDAO struct {
	Driver     string
	DataSource string
}

var db *sql.DB

// ConnectAndCreateTable function launches a new postgres Db instance named go_cars_api and create cars table if not exist
func ConnectAndCreateTable() {
	var (
		err error
		c   = carsDAO{Driver: "postgres", DataSource: "user=kokou password=kokou dbname=go_cars_api sslmode=disable"}
	)

	db, err = sql.Open(c.Driver, c.DataSource)

	errorCheck(err)

	// Create table cars if not exists in postgres Db instance named go_cars_api
	createCarsTable()
}

func createCarsTable() {
	createSQLStatement := "CREATE TABLE IF NOT EXISTS cars(id serial, manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))"

	_, err := db.Exec(createSQLStatement)

	errorCheck(err)
}

// Insert function inserts car right into the table cars inside the postgres Db instance named go_cars_api
func Insert(car *model.Car) {
	// Check if car object exists or not
	objectCheck(car)

	car.CreatedAt = time.Now()
	car.UpdatedAt = time.Now()

	err := GetDbInstance().QueryRow("INSERT INTO cars (manufacturer, design, style, doors, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;", car.Manufacturer, car.Design, car.Style, car.Doors, car.CreatedAt, car.UpdatedAt).Scan(&car.ID)

	errorCheck(err)

	// err := db.Close()

	// errorCheck(err)
}

// FindByID function finds car right from the table cars inside the postgres Db instance named go_cars_api based on ID
func FindByID(id int) *model.Car {
	var (
		car model.Car
	)
	row := GetDbInstance().QueryRow("SELECT * FROM cars WHERE id = $1;", id)
	err := row.Scan(&car.ID, &car.Manufacturer, &car.Design, &car.Style, &car.Doors, &car.CreatedAt, &car.UpdatedAt)

	errorCheck(err)

	return &car
}

// FindAll function finds all cars right from the table cars inside the postgres Db instance named go_cars_api
func FindAll() *model.Cars {
	var (
		cars model.Cars
	)
	rows, err := GetDbInstance().Query("SELECT * FROM cars")

	errorCheck(err)

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var (
			car model.Car
		)

		err := rows.Scan(&car.ID, &car.Manufacturer, &car.Design, &car.Style, &car.Doors, &car.CreatedAt, &car.UpdatedAt)

		errorCheck(err)

		cars = append(cars, car)
	}

	return &cars
}

// Update function updates car right from the table cars inside the postgres Db instance named go_cars_api
func Update(car *model.Car) {

	car.UpdatedAt = time.Now()

	sqlStatement, err := GetDbInstance().Prepare("UPDATE cars SET manufacturer=$1, design=$2, style=$3, doors=$4, updated_at=$5 WHERE id=$6;")

	errorCheck(err)

	_, err = sqlStatement.Exec(car.Manufacturer, car.Design, car.Style, car.Doors, car.UpdatedAt, car.ID)

	errorCheck(err)
}

// DeleteByID function deletes car right from the table cars inside the postgres Db instance named go_cars_api based on ID
func DeleteByID(id int) error {

	sqlStatement, err := GetDbInstance().Prepare("DELETE FROM cars WHERE id=$1;")

	errorCheck(err)

	_, err = sqlStatement.Exec(id)

	return err
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func objectCheck(object interface{}) {
	if object == nil {
		log.Fatal(object)
	}
}

// GetDbInstance function gets postgres Db instance named go_cars_api
func GetDbInstance() *sql.DB {
	return db
}

package main

import (
	"fmt"
	"site-checker/server"
	"site-checker/server/db"

	"github.com/jinzhu/gorm"
)

var connectionString = "host=localhost user=postgres dbname=checker sslmode=disable password=postgres"

func main() {
	var err error
	// init connection to the db
	db.DBCon, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("Database connection error: %s", err.Error()))
	}
	// close connection after stopping
	defer db.DBCon.Close()

	// apply auto migrations to all required models
	// db.DBCon.AutoMigrate(&models.User{}, &models.Lesson{}, &models.Course{})

	// run the web server
	server.RunServer()
}

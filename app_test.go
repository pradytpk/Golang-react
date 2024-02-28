package main

import (
	"database/sql"
	"log"
	"testing"

	"github.com/gorilla/mux"
)

var a App

func TestMain(m *testing.M) {
	app := App{
		Router: &mux.Router{},
		DB:     &sql.DB{},
	}
	err := app.Initialise(DbUser, DbUserPassword, "test")
	if err != nil {
		log.Fatal("Error occured while initalising the database")
	}
	m.Run()
}

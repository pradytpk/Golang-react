package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func main() {
	app := App{
		Router: &mux.Router{},
		DB:     &sql.DB{},
	}
	app.Initialise(DbUser, DbUserPassword, DbName)
	app.Run("localhost:10000")
}

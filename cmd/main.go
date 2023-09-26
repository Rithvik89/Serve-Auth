package main

import (
	"database/sql"
	"net/http"
)

type App struct {
	DB *sql.DB
}

func main() {

	app := &App{}
	initDB(app)
	r := initHandler()

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		panic("Failed in starting server")
	}
}

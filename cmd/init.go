package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rithvik89/auth/utils"
)

var (
	HOST     = utils.GetDBHost()
	PORT     = utils.GetDBPort()
	username = utils.GetDBUsername()
	password = utils.GetDBPassword()
	database = utils.GetDB()
)

func initDB(app *App) {

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	app.DB = db
}

func initKV(app *App) {
	// kv :=

}

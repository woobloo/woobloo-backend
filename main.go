package main

import (
	"database/sql"
	"github.com/gragas/woobloo-backend/config"
	_ "github.com/lib/pq"
	"net/http"
)

var db *sql.DB

func main() {
	// connect to the database
	// TODO: decide on a databse
	// currently using: postgres
	var err error
	config := config.LoadConfig("config.json")
	db, err = sql.Open(config.DBType, config.ConnectionString())
	if err != nil {
		panic(err)
	}

	// setup handlers
	// http.HandleFunc("/", indexHandler)

	http.ListenAndServe(config.ServerAddr, nil)
}

// Handlers

// func indexHandler(w http.ResponseWriter, r *http.Request) {
//	
// }

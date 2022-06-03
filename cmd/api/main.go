package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SolBaa/authservice/data"
)

const webPORT = "8080"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth server on port", webPORT)
	//TODO congif DB

	app := Config{}

	srv := &http.Server{
		Addr:    ":" + webPORT,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

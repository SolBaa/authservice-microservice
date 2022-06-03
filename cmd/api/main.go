package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SolBaa/authservice/data"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPORT = "8080"

var count int

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth server on port", webPORT)
	// congif DB
	conn := connectToDB()
	if conn == nil {
		log.Println("Failed to connect to database. Exiting...")
		os.Exit(1)
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    ":" + webPORT,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Failed to connect to database. Retrying...")
			count++

		} else {
			log.Println("Connected to database")
			return connection
		}
		if count > 10 {
			log.Println("Failed to connect to database. Exiting...")
			os.Exit(1)
		}
		log.Println("Retrying...")
		time.Sleep(time.Second * 2)
		continue
	}
}

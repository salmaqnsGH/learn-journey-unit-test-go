package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
)

type application struct {
	DSN     string
	DB      *sql.DB
	Session *scs.SessionManager
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// set up an app config
	app := application{}

	dsn := os.Getenv("POSTGRES_DSN")
	flag.StringVar(&app.DSN, "dsn", dsn, "Postgres connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = conn

	// get a session manager
	app.Session = getSession()

	// print out message
	log.Println("Starting server on port 8080...")

	// start the server
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

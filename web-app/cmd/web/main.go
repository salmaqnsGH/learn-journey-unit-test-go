package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/salmaqnsGH/unit-test/pkg/db"
)

type application struct {
	DSN     string
	DB      db.PostgresConn
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

	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

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

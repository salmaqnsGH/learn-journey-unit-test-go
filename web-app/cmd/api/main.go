package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/salmaqnsGH/unit-test/pkg/repository"
	"github.com/salmaqnsGH/unit-test/pkg/repository/dbrepo"
)

type application struct {
	DSN       string
	DB        repository.DatabaseRepo
	Domain    string
	JWTSecret string
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var app application

	port := os.Getenv("API_PORT")
	dsn := os.Getenv("POSTGRES_DSN")
	domain := os.Getenv("API_DOMAIN")
	jwtSecret := os.Getenv("API_JWT_SECRET")

	flag.StringVar(&app.DSN, "dsn", dsn, "domain for application, e.g. company.com")
	flag.StringVar(&app.Domain, "domain", domain, "postgres connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", jwtSecret, "signing secret key")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	log.Printf("Starting api on port %s\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

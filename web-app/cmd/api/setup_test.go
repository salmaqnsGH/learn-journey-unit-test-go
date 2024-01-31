package main

import (
	"os"
	"testing"

	"github.com/salmaqnsGH/unit-test/pkg/repository/dbrepo"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE3MDYxNjE3MTEsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.nBdIK2J2PQN0FjR4Gy4w3NSBPWXdly6TWWQSgeqF8U8"

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "SECRet"
	os.Exit(m.Run())
}

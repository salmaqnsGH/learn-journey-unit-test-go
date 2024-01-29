package main

import (
	"os"
	"testing"

	"github.com/salmaqnsGH/unit-test/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "SECRet"
	os.Exit(m.Run())
}

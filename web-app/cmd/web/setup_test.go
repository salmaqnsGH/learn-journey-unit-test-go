package main

import (
	"os"
	"testing"

	"github.com/salmaqnsGH/unit-test/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()

	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}

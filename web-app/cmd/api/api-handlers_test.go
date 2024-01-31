package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/salmaqnsGH/unit-test/pkg/data"
)

func Test_app_authenticate(t *testing.T) {
	var theTests = []struct {
		name               string
		requestBody        string
		expectedStatusCode int
	}{
		{"valid user", `{"email":"admin@example.com","password":"secret"}`, http.StatusOK},
		{"not json", `i'm not json`, http.StatusUnauthorized},
		{"empty json", `{}`, http.StatusUnauthorized},
		{"empty email", `{"email":"","password":"secret"}`, http.StatusUnauthorized},
		{"empty password", `{"email":"admin@example.com","password":""}`, http.StatusUnauthorized},
	}

	for _, e := range theTests {
		reader := strings.NewReader(e.requestBody)
		req, _ := http.NewRequest("POST", "/auth", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)

		handler.ServeHTTP(rr, req)

		if e.expectedStatusCode != rr.Code {
			t.Errorf("%s: returned wrong status code; expected %d but got %d", e.name, e.expectedStatusCode, rr.Code)
		}
	}
}

func Test_app_refresh(t *testing.T) {
	var tests = []struct {
		name               string
		token              string
		expectedStatusCode int
		resetRefreshTime   bool
	}{
		{"valid", "", http.StatusOK, true},
		{"valid but not yet ready to expire", "", http.StatusTooEarly, false},
		{"expired token", expiredToken, http.StatusBadRequest, false},
	}

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}

	oldRefreshTime := refreshTokenExpiry

	for _, test := range tests {
		var token string
		if test.token == "" {
			if test.resetRefreshTime {
				refreshTokenExpiry = time.Second * 1
			}
			tokens, _ := app.generateTokenPair(&testUser)
			token = tokens.RefreshToken
		} else {
			token = test.token
		}

		postedData := url.Values{
			"refresh_token": {token},
		}

		req, _ := http.NewRequest("POST", "/refresh-token", strings.NewReader(postedData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(app.refresh)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedStatusCode {
			t.Errorf("%s: expected status of %d but got %d", test.name, test.expectedStatusCode, rr.Code)
		}
		refreshTokenExpiry = oldRefreshTime
	}
}

func Test_app_userHandlers(t *testing.T) {
	var tests = []struct {
		name           string
		method         string
		json           string
		paramID        string
		handler        http.HandlerFunc
		expectedStatus int
	}{
		{"allUsers", "GET", "", "", app.allUsers, http.StatusOK},
		{"deleteUser", "DELETE", "", "1", app.deleteUser, http.StatusNoContent},
		{"deleteUser bad URL param", "DELETE", "", "A", app.deleteUser, http.StatusBadRequest},
		{"geteUser valid", "GET", "", "1", app.getUser, http.StatusOK},
		{"geteUser bad URL param", "GET", "", "A", app.getUser, http.StatusBadRequest},
		{"geteUser valid", "GET", "", "100", app.getUser, http.StatusBadRequest},
		{
			"updateUser valid",
			"PATCH",
			`{"id":1,"first_name":"John","last_name":"Leni","email":"admin@example.com"}`,
			"",
			app.updateUser,
			http.StatusNoContent,
		},
		{
			"updateUser invalid",
			"PATCH",
			`{"id":100,"first_name":"John","last_name":"Leni","email":"admin@example.com"}`,
			"",
			app.updateUser,
			http.StatusBadRequest,
		},
		{
			"updateUser invalid json",
			"PATCH",
			`"id":1,"first_name":"John","last_name":"Leni","email":"admin@example.com"`,
			"",
			app.updateUser,
			http.StatusBadRequest,
		},
		{
			"insertUser valid",
			"POST",
			`{"id":1,"first_name":"Jack","last_name":"Smith","email":"jack@example.com"}`,
			"",
			app.updateUser,
			http.StatusNoContent,
		},
		{
			"insertUser invalid",
			"POST",
			`{"foo":1}`,
			"",
			app.updateUser,
			http.StatusBadRequest,
		},
		{
			"insertUser invalid json",
			"POST",
			`"id":1,"first_name":"Jack","last_name":"Smith","email":"jack@example.com"`,
			"",
			app.updateUser,
			http.StatusBadRequest,
		},
	}

	for _, e := range tests {
		var req *http.Request
		if e.json == "" {
			req, _ = http.NewRequest(e.method, "/", nil)
		} else {
			req, _ = http.NewRequest(e.method, "/", strings.NewReader(e.json))
		}

		if e.paramID != "" {
			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("userID", e.paramID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(e.handler)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("%s: wrong status returned; expected %d but got %d", e.name, e.expectedStatus, rr.Code)
		}
	}
}

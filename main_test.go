package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var a main.App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestPostMethod(t *testing.T) {
	req, _ := http.NewRequest("POST", "/Team", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetMethod(t *testing.T) {
	req, _ := http.NewRequest("GET", "/Team/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

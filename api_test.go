package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestGet(t *testing.T) {
	a = App{}
	a.Initialize()
	req, _ := http.NewRequest("GET", "/notes", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected not empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

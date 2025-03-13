package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cep=73031031", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, but got %v", status)
	}

	if !strings.Contains(response.Body.String(), `"temp_C"`) {
		t.Errorf("Expected 'temp_C', but got %v", response.Body.String())
	}
}

func TestHandlerInvalidCep(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cep=12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Expected status 422, but got %v", status)
	}

	expected := "invalid zipcode"
	if !strings.Contains(response.Body.String(), expected) {
		t.Errorf("Expected '%s', but got %v", expected, response.Body.String())
	}
}

func TestHandlerNotFoundCep(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cep=99999999", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusNotFound {
		t.Errorf("Expected status 404, but got %v", status)
	}

	expected := "cannot find zipcode"
	if !strings.Contains(response.Body.String(), expected) {
		t.Errorf("Expected '%s', but got %v", expected, response.Body.String())
	}
}

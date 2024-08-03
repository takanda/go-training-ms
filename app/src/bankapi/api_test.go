package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)

func TestStatementHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/statement", nil)
	
	q := req.URL.Query()
	q.Add("number", "1001")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statement)
	
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "1001 - John - 0"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDepositHandler(t *testing.T) {
	jsonBody := []byte(`{"number": 1001, "amount": 10}`)
	req, _ := http.NewRequest("POST", "/statement", bytes.NewBuffer(jsonBody))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deposit)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if accounts[1001].Balance != 10 {
		t.Errorf("deposit could not be executed: result %v want %v", accounts[1001].Balance, 10)
	}
}
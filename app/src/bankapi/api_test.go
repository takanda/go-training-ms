package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/msft/bank"
)

func setup() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
		Balance: 50,
	}
}

func TestStatementHandler(t *testing.T) {
	setup()
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

	expected := "1001 - John - 50"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDepositHandler(t *testing.T) {
	setup()
	jsonBody := []byte(`{"number": 1001, "amount": 10}`)
	req, _ := http.NewRequest("POST", "/statement", bytes.NewBuffer(jsonBody))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deposit)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if accounts[1001].Balance != 60 {
		t.Errorf("deposit could not be executed: result %v want %v", accounts[1001].Balance, 60)
	}
}

func TestWithdrawHandler(t *testing.T) {
	setup()
	jsonBody := []byte(`{"number": 1001, "amount": 5}`)
	req, _ := http.NewRequest("POST", "/withdraw", bytes.NewBuffer(jsonBody))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(withdraw)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if accounts[1001].Balance != 45 {
		t.Errorf("withdraw could not be executed: result %v want %v", accounts[1001].Balance, 45)
	}
}
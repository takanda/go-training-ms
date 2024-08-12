package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/msft/bank"
)

func setup() {
	accounts[1001] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number:  1001,
			Balance: 50,
		},
	}
	accounts[1002] = &CustomAccount{
		Account: &bank.Account{
			Customer: bank.Customer{
				Name:    "Tom",
				Address: "Miami, Florida",
				Phone:   "(207) 333 3119",
			},
			Number:  1002,
			Balance: 75,
		},
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

func TestTransferHandler(t *testing.T) {
	setup()
	jsonBody := []byte(`{"number1": 1001, "number2": 1002, "amount": 15}`)
	req, _ := http.NewRequest("POST", "/transfer", bytes.NewBuffer(jsonBody))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transfer)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}
	
	if accounts[1001].Balance != 35 || accounts[1002].Balance != 90 {
		t.Errorf("transfer could not be executed correctly: result 1001-%v, 1002-%v want 1001-%v, 1002-%v",
		          accounts[1001].Balance, accounts[1002].Balance, 35, 90)
	}
}
package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
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
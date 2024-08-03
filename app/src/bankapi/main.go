package main

import (
	"github.com/msft/bank"
	"net/http"
	"fmt"
	"strconv"
	"log"
	"io"
	"encoding/json"
)

var accounts = map[int32]*bank.Account{}

func init() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}
}

func main() {
	http.HandleFunc("GET /statement", statement)
	http.HandleFunc("POST /deposit", deposit)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	number, err := strconv.Atoi(numberqs)
	if err != nil {
		fmt.Fprintf(w, "Invalid account number")
		return
	}

	number32 := int32(number)
	if err != nil {
		fmt.Fprintf(w, "Invalid account number")
	} else {
		if account, ok := accounts[number32]; !ok {
			fmt.Fprintf(w, "Account with number %v can't be found", number32)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	defer req.Body.Close()

	var requestData map[string]interface{}
	if err := json.Unmarshal(body, &requestData); err != nil {
		fmt.Fprintf(w, "Invalid JSON format")
		return
	}
	
	number64, ok := requestData["number"].(float64)
	number32 := int32(number64)
	if !ok {
		fmt.Fprintf(w, "Invalid number")
		return
	}
	amount, ok := requestData["amount"].(float64)
	if !ok {
		fmt.Fprintf(w, "Invalid amount")
		return
	}

	if account, ok := accounts[number32]; !ok {
		fmt.Fprintf(w, "Account with number %v can't be found", number32)
	} else {
		err := account.Deposit(amount)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}
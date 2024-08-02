package main

import (
	"github.com/msft/bank"
	"net/http"
	"fmt"
	"strconv"
	"log"
)

var accounts = map[float64]*bank.Account{}

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
	http.HandleFunc("/statement", statement)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number")
	} else {
		account, presence := accounts[number]

		if presence {
			fmt.Fprintf(w, account.Statement())
		} else {
			fmt.Fprintf(w, "Account with number %v can't be found", number)
		}
	}
}
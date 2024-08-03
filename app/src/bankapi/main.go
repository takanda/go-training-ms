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

var accounts = map[int32]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &bank.Account {
			Customer: bank.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}
	accounts[1002] = &CustomAccount{
		Account: &bank.Account {
			Customer: bank.Customer{
				Name:    "Tom",
				Address: "Miami, Florida",
				Phone:   "(207) 333 3119",
			},
			Number: 1002,
			Balance: 75,
		},
	}
	http.HandleFunc("GET /statement", statement)
	http.HandleFunc("POST /deposit", deposit)
	http.HandleFunc("POST /withdraw", withdraw)

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
			json.NewEncoder(w).Encode(bank.Statement(account))
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

func withdraw(w http.ResponseWriter, req *http.Request) {
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

	floatNumber, ok := requestData["number"].(float64)
	intNumber := int32(floatNumber)
	if !ok {
		fmt.Fprintf(w, "Invalid number")
		return
	}
	
	amount, ok := requestData["amount"].(float64)
	if !ok {
		fmt.Fprintf(w, "Invalid amount")
		return
	}

	if account, ok := accounts[intNumber]; !ok {
		fmt.Fprintf(w, "Account with number %v can't be found", intNumber)
	} else {
		err := account.Withdraw(amount)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	// A → B
	// 1. A - amount ・・・withdraw
	// 2. B + amount ・・・deposit
	// Transaction needed

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	defer req.Body.Close()
	
	var requestData map[string]interface{}
	// requestDataを渡す場合、Unmarshalの第二引数は上のrequestDataとは別のメモリに格納される変数(中身はrequestDataのコピー)
	// しかし、後続処理で使用するrequestDataは上のrequestDataが格納されたメモリを参照する
	// そのため、Unmarshalでbodyを受け取っていないrequestDataを参照するため後続処理が正常に作動しない
	// なのでUnmarshalにはポインターを渡す必要がある
	// Unmarshalにはnilかポインターを渡す必要があるが、上のrequestDataの定義の段階ではnilなので
	// requestDataを渡してもエラーにはならない
	if err := json.Unmarshal(body, &requestData); err != nil {
		fmt.Fprintf(w, "Invalid JSON")
		return
	}

	float64Number1, ok := requestData["number1"].(float64)
	int32Number1 := int32(float64Number1)
	if !ok {
		fmt.Fprintf(w, "Invalid number")
		return
	}
	float64Number2, ok := requestData["number2"].(float64)
	int32Number2 := int32(float64Number2)
	if !ok {
		fmt.Fprintf(w, "Invalid number")
		return
	}
	
	amount, ok := requestData["amount"].(float64)
	if !ok {
		fmt.Fprintf(w, "Invalid amount")
		return
	}

	if account1, ok := accounts[int32Number1]; !ok {
		fmt.Fprintf(w, "Account with number %v can't be found", int32Number1)
		return
	} else {
		err := account1.Withdraw(amount)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		}
	}
	
	if account2, ok := accounts[int32Number2]; !ok {
		fmt.Fprintf(w, "Account with number %v can't be found", int32Number1)
		return
	} else {
		err := account2.Deposit(amount)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}


type CustomAccount struct {
	*bank.Account
}

func (c *CustomAccount) Statement() string {
	json, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(json)
}
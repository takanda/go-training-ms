package basic

import (
	"fmt"
	// "os"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func BasicErrorHandling() {
	employee, err := getInformation(1001)
    if err != nil {
        // Something is wrong. Do something.
    } else {
        fmt.Println(employee)
    }
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	employee.Address = "Street"
	return employee, err
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

package structure

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	// Information Person
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func LearnStructure() {
	// john := Employee{1001, "John", "Doe", "Doe's Street"}
	// john := Employee{FirstName: "John", LastName: "Doe"}
	// john.ID = 1001
	// john.FirstName = "Bob"
	// employeeCopy := &john
	// employeeCopy.Address = "Street!"
	// fmt.Println(john)
	EmbedStructure()
}

func EmbedStructure() {
	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Doe"
	fmt.Println(employee)
}


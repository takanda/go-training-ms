package structure

import "fmt"
import "encoding/json"

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

	// EmbedStructure()
	JsonStructure()
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

type Person2 struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee2 struct {
	Person2
	ManagerID int
}

type Contractor2 struct {
	Person2
	CompanyID int
}


func JsonStructure() {
	employees := []Employee2{
		Employee2{
			Person2: Person2{
				LastName: "Doe", FirstName: "Joe",
			},
		},
		Employee2{
			Person2: Person2{
				LastName: "Cambell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee2
    json.Unmarshal(data, &decoded)
    fmt.Printf("%v", decoded)
}

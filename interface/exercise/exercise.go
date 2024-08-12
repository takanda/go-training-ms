package exercise

import "fmt"

type Account struct {
	FirstName, LastName string
}

func (a Account) ChangeName(F string, L string) {
	a.FirstName = F
	a.LastName = L
}

func (a Account) String() string {
	return fmt.Sprintf("First name is %s, and last name is %s", a.FirstName, a.LastName)
}

type Employee struct {
	Account
	credit float64
}

func (e *Employee) AddCredits(c float64) {
	e.credit += c
}
func (e *Employee) RemoveCredits(c float64) {
	e.credit -= c
}
func (e *Employee) CheckCredits() float64 {
	return e.credit
}

func RunExercise() {
	person1 := Account{FirstName: "John", LastName: "Doe"}
	person2 := Account{FirstName: "Michael", LastName: "Smith"}
	fmt.Printf("%s\n%s\n", person1, person2)

	employee1 := &Employee{
		Account: Account{FirstName: "Paul", LastName: "A"},
		credit: 10,
	}
	fmt.Println("Now, credit is,", employee1.CheckCredits())
	employee1.AddCredits(20.3)
	fmt.Println("Now, credit is,", employee1.CheckCredits())
}


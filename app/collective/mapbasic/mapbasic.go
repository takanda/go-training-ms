package mapbasic

import "fmt"

func LearnMap() {
	// AddMap()
	// AccessMap()
	// DeleteMap()
	LoopMap()
}

func AddMap() map[string]int {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 21
	return studentsAge
}

func AccessMap() {
	studentsAge := AddMap()

	age, exist := studentsAge["john"]
	if exist {
		fmt.Println("John's age is", age)
	} else {
		fmt.Println("John's age couldn't be found")
	}
}

func DeleteMap() {
	studentsAge := AddMap()
	delete(studentsAge, "bob")
	fmt.Println(studentsAge)
}

func LoopMap() {
	studentsAge := AddMap()
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
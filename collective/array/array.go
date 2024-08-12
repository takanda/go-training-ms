package array

import "fmt"

func LearnArray() {
	var a [3]int
	a[1] = 10
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
    fmt.Println("Cities:", cities)
	
	cities2 := [...]string{"New York", "Paris", "Berlin", "Madrid"}
    fmt.Println("Cities:", cities2)

	numbers := [...]int{99: -1}
	fmt.Println("First Position:",numbers[0])
	fmt.Println("Last Position:",numbers[99])
	fmt.Println("Length:",len(numbers))

	var twoD [3][5]int
	fmt.Println("initial status:", twoD)
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i+1) * (j+1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)
}
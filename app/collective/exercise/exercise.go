package exercise

import (
	"fmt"
)

func RunCalcFib() {
	var number int
	fmt.Print("Please enter number:")
	fmt.Scanf("%d", &number)
	fmt.Println("Result:", CalcFib(number))
}

func CalcFib(number int) []int {
	result := make([]int, number)
	defer func() {
		handler := recover()
		if handler != nil {
			result = make([]int, 0)
		}
	}()
	if number <= 2 {
		panic("Input must be greater than 2")
	}

	result[0], result[1] = 1, 1

	for i := 2; i < number; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result
}

package exercise

import (
	"fmt"
)


func CalcFib() {
	var number int
	var result []int
	fmt.Print("Please enter number:")
	fmt.Scanf("%d", &number)
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("Error", result)
		}
	}()
	if number <= 2 {
		panic("Input must be greater than 2")
	}

	for i := 1; i <= number; i++ {
		if i <= 2 {
			result = append(result, 1)
			continue
		}
		
		prevNumbers := result[len(result)-2:]
		appendNumber := 0
		for _, num := range prevNumbers {
			appendNumber += num
		}
		result = append(result, appendNumber)
	}
	fmt.Println("Result is:", result)
}
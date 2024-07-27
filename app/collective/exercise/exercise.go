package exercise

import (
	"fmt"
)

func RunCalcFib() {
	var number int
	fmt.Print("Please enter number:")
	fmt.Scanf("%d", &number)
	fmt.Println("Result:", calcFib(number))
}

func calcFib(number int) []int {
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

func RunAlphaToNum() {
	fmt.Println("MCLX is", alphaToNum("MCLX"))
	fmt.Println("MCMXCIX is ", alphaToNum("MCMXCIX"))
	fmt.Println("MCMZ is", alphaToNum("MCMZ"))
}

func alphaToNum(alpha string) int {
	alphaMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	alphaSlice := make([]int, len(alpha)+1)
	result := 0

	defer func() {
		handler := recover()
		if handler != nil {
			result = 0
		}
	}()

	for index, value := range alpha {
		if val, present := alphaMap[value]; present {
			alphaSlice[index] = val
		} else {
			panic("Error")
		}
	}

	for i := 0; i < len(alpha); i++ {
		if alphaSlice[i] < alphaSlice[i+1] {
			result -= alphaSlice[i]
		} else {
			result += alphaSlice[i]
		}
	}

	return result
}
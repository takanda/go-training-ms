package exercise

import (
	"fmt"
	"math"
)

func fib(target int) []int {
	result := make([]int, target)
	result[0], result[1] = 1, 1

	if target <= 2 {
		result = make([]int, 0)
		return result
	}

	for i := 2; i < target; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result
}

func RunFib() {
	var target int
	fmt.Print("Please enter number: ")
	fmt.Scanf("%d", &target)
	fmt.Println(fib(target))
}

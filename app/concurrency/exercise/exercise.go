package exercise

import (
	"fmt"
	"time"
	"math/rand"
)

func fib(target int, ch chan string) {
	result := make([]int, target)
	result[0], result[1] = 1, 1

	if target <= 2 {
		result = make([]int, 0)
		ch <- fmt.Sprintf("%v\n", result) 
	}

	for i := 2; i < target; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	r := rand.Intn(3)
    time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("%v\n", result)
}

func RunFib() {
	start := time.Now()

	ch := make(chan string)

	var target int
	fmt.Print("Please enter number: ")
	fmt.Scanf("%d", &target)
	for i := 0; i < target; i++ {
		go fib(target, ch)
	}

	for i := 0; i < target; i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

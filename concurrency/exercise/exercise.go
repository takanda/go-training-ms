package exercise

import (
	"fmt"
	"time"
	"math/rand"
)

func fib(target int, ch chan int) {
	x, y := 1, 1
	for i := 0; i < target; i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
    time.Sleep(time.Duration(r) * time.Second)

	ch <- x
}

func RunFib() {
	target := UserInterface()
	start := time.Now()
	
	ch := make(chan int, target)
	for i := 0; i < target; i++ {
		go fib(i, ch)
	}
	for i := 0; i < target; i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

var quit = make(chan bool)

func fib2(c chan int) {
	x, y := 1, 1

	for {
		select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("Done calculating Fibonacci!")
				return
		}
	}
}

func RunFib2() {
	start := time.Now()
	command := ""
	data := make(chan int)

	go fib2(data)

    for {
        num := <-data
        fmt.Println(num)
        fmt.Scanf("%s", &command)
        if command == "quit" {
            quit <- true
            break
        }
    }

	time.Sleep(1 * time.Second)

    elapsed := time.Since(start)
    fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func UserInterface() int {
	var target int
	fmt.Print("Please enter number: ")
	fmt.Scanf("%d", &target)
	return target
}
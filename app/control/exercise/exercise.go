package exercise

import "fmt"

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch  {
		case i % (3*5) == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func FindPrimeNumber(number int) {
	for i := 1; i <= number; i++ {
		result := true
		if i == 1 {
			continue
		}

		for j := 2; j < i; j++ {
			if i % j == 0 {
				result = false
				break
			}
		}

		if result {
			fmt.Printf("%v\n", i)
		}
	}
}

func NegativeNumBecomePanic() {
	var val int
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val == 0:
			fmt.Println(val, "is neither negative nor positive")
		case val > 0:
			fmt.Println("You entered:", val)
		case val < 0:
			panic("negative value")
	}
	}
}
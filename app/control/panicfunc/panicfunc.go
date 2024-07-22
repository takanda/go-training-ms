package panicfunc

import "fmt"

func Panicfunc(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("higlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
    fmt.Println("Call: highlow(", high, ",", low, ")")

    Panicfunc(high, low + 1)
}
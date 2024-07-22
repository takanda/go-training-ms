package recoverfunc

import (
	"fmt"
	"control/panicfunc"
)


func Recoverfunc(high int, low int) {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	panicfunc.Panicfunc(high, low)
}
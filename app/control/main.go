package main

import "fmt"
import "control/panicfunc"



func main() {
	panicfunc.Panicfunc(2, 0)
	fmt.Println("Program finished successfully!")
}
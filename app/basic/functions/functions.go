package functions

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )
import "fmt"

func ShowFunctions() {
	// sum, mul := calc(os.Args[1], os.Args[2])
	// fmt.Println("Sum:", sum)
	// fmt.Println("Mul:", mul)
	firstName := "John"
	fmt.Println(firstName)
	updateName(&firstName)
	fmt.Println(firstName)
}

// func calc(number1 string, number2 string) (sum int, mul int) {
// 	int1, _ := strconv.Atoi(number1)
// 	int2, _ := strconv.Atoi(number2)
// 	sum = int1 + int2
// 	mul = int1 * int2
// 	return 
// }

// ポインタ変数への値の割り当てはデリファレンス(*をつける)しなければエラーになる
// structのフィールドにアクセスする時のみ自動的なデリファレンスが実行される
func updateName(name *string) {
	*name = "David"
}
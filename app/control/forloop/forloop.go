package forloop

import "math/rand"
import "time"

func Forloop() (int64, int) {
	// sum := 0
	// for i := 0; i <= 100; i+=10 {
	// 	sum += i
	// }
	// return sum

	var num int64
	try_count := 0
	rand.Seed(time.Now().Unix())

	for num != 5 {
		num = rand.Int63n(15)
		try_count += 1
	}
	return num, try_count
}
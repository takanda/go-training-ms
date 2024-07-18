package switchstate

import (
	"time"
)

func Switchstate() string {
	// sec := time.Now().Unix()
	// rand.Seed(sec)
	// i := rand.Int31n(10)

	// switch i {
	// case 0:
	// 	return "zero"
	// case 1:
	// 	return "one"
	// case 2:
	// 	return "two"
	// default:
	// 	return "no match"
	// }

	switch time.Now().Weekday().String() {
		case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
			return "Weekday!"
		default:
			return "Weekend!"
	}
}
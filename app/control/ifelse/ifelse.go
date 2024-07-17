package ifelse

import "strconv"

// func Ifelse() string {
// 	x := 27
// 	if x%2 == 0 {
// 		return strconv.Itoa(x) + " is even number."
// 	}
// 	return strconv.Itoa(x) + " is odd number."
// }

func givemeagnumber() int {
	return -1
}

func Ifelse() string {
	if num := givemeagnumber(); num < 0 {
		return strconv.Itoa(num) + " is negative."
	} else if num < 10 {
		return strconv.Itoa(num) + " has only one digit."
	} else {
		return strconv.Itoa(num) + " has multiple digits."
	}
}
package switchstate

import (
	"regexp"
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

	// switch time.Now().Weekday().String() {
	// 	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 		return "Weekday!"
	// 	default:
	// 		return "Weekend!"
	// }
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
    var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

    contact := "foo@bar.com"

    switch {
    case email.MatchString(contact):
        return("email")
    case phone.MatchString(contact):
        return("phone number")
    default:
        return("not match")
    }
}
package slice

import "fmt"

func LearnSlice() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
    fmt.Println(months)
    quarter1 := months[0:3]
    quarter2 := months[3:6]
    quarter3 := months[6:9]
    quarter4 := months[9:12]
    fmt.Println(quarter1, len(quarter1), cap(quarter1))
    fmt.Println(quarter2, len(quarter2), cap(quarter2))
    fmt.Println(quarter3, len(quarter3), cap(quarter3))
    fmt.Println(quarter4, len(quarter4), cap(quarter4))

	fmt.Println(quarter2[2])
}
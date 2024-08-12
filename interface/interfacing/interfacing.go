package interfacing

import "fmt"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s *Square) Area() float64 {
    return s.size * s.size
}

func (s *Square) Perimeter() float64 {
	return s.size * 4
}


func LearnInterface() {
	// var s Shape = &Square{3}
	s := Shape(&Square{3})
	fmt.Println(s.Area())
}

type Person struct {
    Name, Country string
}

func (p Person) String() string {
    return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}
func StringerInterface() {
    rs := Person{"John Doe", "USA"}
    ab := Person{"Mark Collins", "United Kingdom"}
    fmt.Printf("%s\n%s\n", rs, ab)
}

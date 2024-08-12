package method

import "fmt"

type triangle struct {
	size int
}

type square struct {
    size int
}

type colorTriangle struct {
	triangle
	color string
}

func (t *triangle) perimeter() int {
	return t.size * 3
}

func (s *square) perimeter() int {
    return s.size * 4
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func (t *colorTriangle) perimeter() int {
	return t.size * 3 * 2
}

func LearnMethod() {
    t := triangle{size: 3}
    t.doubleSize()
    fmt.Println("Size:", t.size)
    fmt.Println("Perimeter:", t.perimeter())
}


func EmbedMethod() {
	t := colorTriangle{triangle{4}, "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter(colored):", t.perimeter())
	fmt.Println("Perimeter(normal):", t.triangle.perimeter())
}
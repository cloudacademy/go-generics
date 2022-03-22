package main

import (
	"fmt"
)

type number interface {
	int | float64
}

func area[T number](a, b T) T {
	return a * b
}

func main() {
	fmt.Println(area(10, 20))
	fmt.Println(area(3.25, 5.75))
}

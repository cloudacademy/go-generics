package main

import (
	"fmt"
	"strings"
)

type StarableInt interface {
	~int
	Stars() string
}

type Rating int

func (r Rating) Stars() string {
	return strings.Repeat("*", int(r))
}

/*
func RatingToStars[T StarableInt](s T) string {
	return s.Stars()
}
*/

func RatingToStars[T interface {
	~int
	Stars() string
}](s T) string {
	return s.Stars()
}

func main() {
	fmt.Println()

	var r1 Rating = 5
	var r2 Rating = 3

	stars1 := RatingToStars(r1)
	stars2 := RatingToStars(r2)

	fmt.Printf("rating %d -> %s\n", r1, stars1)
	fmt.Printf("rating %d -> %s\n", r2, stars2)
}

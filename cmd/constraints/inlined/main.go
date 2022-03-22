package main

import (
	"fmt"
)

/*
type SliceConstraint[E any] interface {
	~[]E
}

func firstElement[S SliceConstraint[E], E any](s S) E {
	return s[0]
}

func lastElement[S SliceConstraint[E], E any](s S) E {
	return s[len(s)-1]
}
*/

func firstElement[S interface{ ~[]E }, E any](s S) E {
	return s[0]
}

func lastElement[S interface{ ~[]E }, E any](s S) E {
	return s[len(s)-1]
}

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println()

	experts := []Person{{"Euler", 27}, {"Newton", 39}, {"Einstein", 25}}

	fmt.Println(firstElement(experts))
	fmt.Println(lastElement(experts))
}

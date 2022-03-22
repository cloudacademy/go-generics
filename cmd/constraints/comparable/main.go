package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func getBiggerNumber1[T numbers](num1, num2 T) T {
	if num1 > num2 {
		return num1
	}
	return num2
}

func getBiggerNumber2[T constraints.Ordered](num1, num2 T) T {
	if num1 > num2 {
		return num1
	}
	return num2
}

func makeComparableFunc[T comparable]() func(lhs, rhs T) bool {
	return func(lhs, rhs T) bool {
		return lhs == rhs
	}
}

func main() {
	fmt.Println()

	fmt.Println(getBiggerNumber1(100, 200))
	fmt.Println(getBiggerNumber2(22.0/7.0, math.Pi))

	areEqual := makeComparableFunc[int]()
	fmt.Println(areEqual(1, 2))
	fmt.Println(areEqual(20, 20))
	fmt.Println(areEqual(len("cloud"), len("academy")))
}

package main

import (
	"fmt"
)

func filter[T any](data []T, f func(T) bool) []T {
	var result []T
	for _, e := range data {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

func main() {
	fmt.Println()

	result1 := filter([]int{1, 2, 3, 4, 5, 6}, func(v int) bool { return v <= 3 })
	fmt.Println(result1)

	result2 := filter([]string{"cloudacademy", "golang", "generics", "fun"}, func(v string) bool { return len(v) > 6 })
	fmt.Println(result2)
}

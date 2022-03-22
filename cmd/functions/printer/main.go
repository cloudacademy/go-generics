package main

import (
	"fmt"
)

func printer[T any](data T) {
	fmt.Println(data)
}

func main() {
	fmt.Println()

	printer("cloudacademy")
	printer(22.0 / 7.0)
	printer(true)
}

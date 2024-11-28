package main

import "fmt"

const (
	_ = iota + 1991
	first
	second
	third
	fourth
)

func main() {
	fmt.Printf("First: %v\nSecond: %v\nThird: %v\nFourth: %v\n", first, second, third, fourth)
}

package main

import "fmt"

type goLearn int

var x goLearn

func main() {
	fmt.Printf("%v,%T\n", x, x)
	x = 42
	fmt.Printf("%v", x)

}

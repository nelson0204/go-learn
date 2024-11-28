package main

import "fmt"

type goLearn int

var x goLearn
var y int

func main() {
	fmt.Printf("%v,%T\n", x, x)
	x = 42
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Printf("%v,%T\n", y, y)
}

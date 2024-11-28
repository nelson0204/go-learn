package main

import (
	"fmt"
)

func main() {
	a := (9 == 9)
	b := (9 != 9)
	c := (9 <= 9)
	d := (9 < 9)
	e := (9 >= 9)
	f := (9 > 9)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}

package main

import "fmt"

func main() {
	x := 100

	fmt.Printf("Decimal: %d\nBinário:%b\nHexadecimal:%#x\n", x, x, x)

	y := x << 1
	fmt.Printf("Decimal: %d\nBinário:%b\nHexadecimal:%#x\n", y, y, y)

}

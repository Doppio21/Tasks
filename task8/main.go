package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func printBits(a int64) {
	for i := int64(1 << 62); i >= 1; i >>= 1 {
		if (a & i) > 0 {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
	fmt.Println()
}

func main() {
	bitNum := 3
	set := false

	a := int64(1678829)
	printBits(a)

	mask := int64((1 << (bitNum - 1)))
	if set {
		a |= mask
	} else {
		a &= ^mask
	}
	
	printBits(a)
}

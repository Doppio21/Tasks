package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 9
	b := 7

	a, b = Swap2(a, b)
	fmt.Println("a =",a, "b =", b)
}

func Swap1(a int, b int) (int, int){
	a,b=b,a
	return a, b
}

func Swap2(a int, b int) (int, int){
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

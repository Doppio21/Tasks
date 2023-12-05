package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	// ввод переменных в аргументах командной строки
	// "a (+,-,*,/) b"
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("input error%v", err)
		return
	}

	sign := os.Args[2]

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("input error%v", err)
		return
	}

	if a > int(math.Pow(2, 20)) && b > int(math.Pow(2, 20)) {
		switch sign {
		case "*":
			fmt.Printf("%d * %d = %d", a, b, a*b)
		case "/":
			fmt.Printf("%d / %d = %d", a, b, a/b)
		case "+":
			fmt.Printf("%d + %d = %d", a, b, a+b)
		case "-":
			fmt.Printf("%d - %d = %d", a, b, a-b)

		}
	} else {
		fmt.Println("value is less 2^20")
	}
}

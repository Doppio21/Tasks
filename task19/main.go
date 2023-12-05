package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func ReverseStr(str string) string {
	runeStr := []rune(str)
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}

func main() {
	inputStr := "главрыба"
	fmt.Println(ReverseStr(inputStr))
}

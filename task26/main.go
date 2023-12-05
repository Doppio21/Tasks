package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func CheckOriginal(str string) bool {
	str = strings.ToLower(str)
	for _, s := range str {
		count := strings.Count(str, string(s))
		if count > 1 {
			return false
		}
	}
	return true
}

func CheckOriginal2(str string) bool {
	c := map[rune]struct{}{}

	for _, r := range []rune(str) {
		lr := unicode.ToLower(r)
		if _, ok := c[lr]; ok {
			return false
		}

		c[lr] = struct{}{}
	}

	return true
}

func main() {
	str := "abCdefAaf"
	if CheckOriginal(str) && CheckOriginal2(str) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

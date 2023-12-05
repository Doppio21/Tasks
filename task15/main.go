package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// При вызове функции createHugeString() создается строка v,
// занимающаяя большое колличество оперативной памяти
// Происходит присваивание переменной justString урезанного значения v,
// при этом строка v остается в памяти.
// Также использование глобальной переменной
// может вызывать race condition при конкурентном вызове someFunc()

// https://stackoverflow.com/a/55860599
func createHugeString(n int) string {
	buff := make([]byte, int(math.Ceil(float64(n)/float64(4/3))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:n]
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return strings.Clone(v[:100])
}

func main() {
	s := someFunc()
	fmt.Println(s)
}

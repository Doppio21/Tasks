package main

import (
	"fmt"
	"os"
	"strconv"
)

// Удалить i-ый элемент из слайса.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("input error%v", err)
		return
	}

	arr = append(arr[:i], arr[i+1:]...)

	// Если пордяок не важен
	// arr[i] = arr[len(arr)-1]
	// arr = arr[:len(arr)-1]

	fmt.Println(arr)
}

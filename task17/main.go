package main

import (
	"fmt"
	"os"
	"strconv"
)

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(arr []int, num int) int {
	for i, j := 0, len(arr)-1; i <= j; {
		mid := (j + i) / 2
		numMid := arr[mid]
		if numMid < num {
			i = mid + 1
		} else if numMid > num {
			j = mid - 1
		} else if numMid == num {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("error entering the number%v", err)
		return
	}
	res := binarySearch(arr, num)
	if res == -1 {
		fmt.Println("not found")
	} else {
		fmt.Println(res)
	}
}

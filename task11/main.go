package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	a := []int{5, 4, 9, 6, 2, 1}
	b := []int{11, 14, 7, 5, 9}
	res := make([]int, 0)

	for _, v := range a {
		for _, v2 := range b {
			if v == v2 {
				res = append(res, v)
			}
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

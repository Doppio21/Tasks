package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	res := make(map[string]struct{})
	for _, key := range a {
		res[key] = struct{}{}
	}
	for k := range res {		
		fmt.Println(k)
	}
}

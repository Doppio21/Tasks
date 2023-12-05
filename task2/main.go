package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func Square(a *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*a = *a * *a
}

func main() {
	wg := sync.WaitGroup{}
	m := []int{2, 4, 6, 8, 10}

	for i := range m {
		wg.Add(1)
		go Square(&m[i], &wg)
	}

	wg.Wait()
	for _, v := range m {
		fmt.Println(v)
	}
}

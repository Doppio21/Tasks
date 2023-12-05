package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func SumSquare(a int, res *atomic.Int32, wg *sync.WaitGroup) {
	defer wg.Done()
	res.Add(int32(a * a))
}

func main() {
	var (
		wg  = sync.WaitGroup{}
		m   = []int{2, 4, 6, 8, 10}
		res atomic.Int32
	)

	for _, v := range m {
		wg.Add(1)
		go SumSquare(v, &res, &wg)
	}
	wg.Wait()

	fmt.Println(res.Load())
}

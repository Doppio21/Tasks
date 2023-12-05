package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	Count atomic.Int32
}

func main() {
	wg := sync.WaitGroup{}
	counter := Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Count.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(int(counter.Count.Load()))
}

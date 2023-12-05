package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.


func main() {
	timeout, err := strconv.Atoi(os.Args[1])
	if err != nil || timeout == 0 {
		fmt.Printf("error entering the number of workers %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case value := <-ch:
				fmt.Println(value)
			}
		}
	}()

loop:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			break loop
		case ch <- i:
		}
	}

	wg.Wait()
}

package main

import (
	"context"
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала:
// в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

type Conveyor[T any] struct {
	funcs    []func(T) T
	channels []chan T
	wg       sync.WaitGroup
}

func NewConveyor[T any](fs []func(T) T) *Conveyor[T] {
	ret := &Conveyor[T]{
		funcs:    fs,
		channels: make([]chan T, len(fs)+1),
	}

	for i := 0; i < len(ret.channels); i++ {
		ret.channels[i] = make(chan T, 1)
	}

	return ret
}

func (c *Conveyor[T]) Run(ctx context.Context) {
	for i, f := range c.funcs {
		c.wg.Add(1)
		go func(i int, f func(T) T) {
			defer c.wg.Done()
			in := c.channels[i]
			out := c.channels[i+1]

			for {
				var v T
				select {
				case <-ctx.Done():
					return
				case v = <-in:
				}

				updated := f(v)

				select {
				case <-ctx.Done():
					return
				case out <- updated:
				}
			}
		}(i, f)
	}
}

func (c *Conveyor[T]) Push(ctx context.Context, v T) {
	select {
	case <-ctx.Done():
		return
	case c.channels[0] <- v:
	}
}

func (c *Conveyor[T]) Pop(ctx context.Context) T {
	select {
	case <-ctx.Done():
		return *new(T)
	case v := <-c.channels[len(c.channels)-1]:
		return v
	}
}

func (c *Conveyor[T]) WaitStop() {
	c.wg.Wait()
}

func main() {
	funcs := []func(a int) int{
		func(a int) int {
			return a * 2
		},
		// func(a int) int {
		// 	return a + 2
		// },
		// func(a int) int {
		// 	return a * a
		// },
	}
	c := NewConveyor[int](funcs)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c.Run(ctx)
	defer c.WaitStop()

	values := []int{1, 2, 3, 4, 5}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, v := range values {
			u := c.Pop(ctx)
			fmt.Println("Pop", u)

			expected := v
			for _, f := range funcs {
				expected = f(expected)
			}

			if u != expected {
				panic("invalid value")
			}
		}
	}()

	for _, v := range values {
		c.Push(ctx, v)
		fmt.Println("Push", v)
	}

	wg.Wait()
	cancel()
}

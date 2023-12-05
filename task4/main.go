package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

type Pool struct {
	ch    chan []byte
	wg    sync.WaitGroup
	count int
}

func NewPool(count int) *Pool {
	return &Pool{
		ch:    make(chan []byte, count),
		count: count,
	}
}

func (w *Pool) Run(ctx context.Context) {
	for i := 0; i < w.count; i++ {
		w.wg.Add(1)
		go func(i int) {
			defer w.wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("worker %d stopped\n", i+1)
					return
				case data := <-w.ch:
					fmt.Println("data from channel:\n", data)
				}
			}
		}(i)
	}
}

func (w *Pool) Stop() {
	w.wg.Wait()
	close(w.ch)
}

func main() {
	// Завершить по сигналу
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Ввод количества воркеров в аргументах командной строки
	count, err := strconv.Atoi(os.Args[1])
	if err != nil || count == 0 {
		fmt.Printf("error entering the number of workers %v", err)
		return
	}

	pool := NewPool(count)
	pool.Run(ctx)
	defer pool.Stop()

	// постоянная запись в канал 
	for {
		data := make([]byte, 16)
		if _, err := rand.Read(data); err != nil {
			fmt.Printf("data generation error %v", err)
			return
		}

		select {
		case <-ctx.Done():
			fmt.Println("the program has been completed")
			return
		case pool.ch <- data:
		}
	}
}

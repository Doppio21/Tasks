package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	// Остановить по сигналу
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Остановить через 20 секунд
	ctxTime, cancelTime := context.WithTimeout(ctx, 20*time.Second)
	defer cancelTime()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctxTime.Done()
		fmt.Println("goroutime stopped")
	}()

	<-ctxTime.Done()

	wg.Wait()
	fmt.Println("exited")
}

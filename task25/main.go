package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func Sleep(t time.Duration) {
	<-time.After(t)
}

func Sleep2(t time.Duration) {
	now := time.Now()
	for time.Since(now) < t {
	}
}

func main() {
	now := time.Now()
	sleep := 1 * time.Second

	Sleep(sleep)
	Sleep2(sleep)

	fmt.Printf("spent %f seconds\n", time.Since(now).Seconds())
}

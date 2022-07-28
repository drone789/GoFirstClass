package main

import (
	"sync"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 10
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for i := range ch {
		println(i)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}

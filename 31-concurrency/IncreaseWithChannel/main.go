package main

import (
	"fmt"
	"sync"
)

type counter struct {
	c chan int
	i int
}

func NewCounter() *counter {
	cter := &counter{
		c: make(chan int),
		i: 0,
	}
	go func() {
		for {
			cter.i++
			// [cter.c 通道中的值没有被取走，会阻塞]
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter) Increase() int {
	// 从无缓冲区的chanel中接受数据后,[cter.c <- cter.i 阻塞得到释放]
	return <-cter.c
}

func main() {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

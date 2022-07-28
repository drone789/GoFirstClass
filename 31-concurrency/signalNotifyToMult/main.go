package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done \n", i)
}

type signal struct{}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)

	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)

		go func(i int) {
			// 3 等待信号来,开始干活
			<-groupSignal
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		// 4 告诉调用者，活干完了
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("select a group of workers...")
	groupSignal := make(chan signal)

	// 1 Goroutine 启动后会阻塞在名为 groupSignal 的无缓冲 channel 上
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(3 * time.Second)
	fmt.Println("the group of workers start to work...")

	// 2 向所有 worker goroutine 广播“开始工作”的信号
	close(groupSignal)

	// 5 阻塞,直到收到干完活的通知
	<-c
	fmt.Println("the group of workers work done!")
}

package main

import "time"

type signal struct{}

func worker() {
	println("woker is working....")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("work start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	println("start a worker....")
	c := spawn(worker)
	// 阻塞 等待spawn执行完，有返回值
	<-c

}

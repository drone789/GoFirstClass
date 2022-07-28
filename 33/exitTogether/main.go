package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 7)
		ch2 <- 7
		close(ch2)
	}()

	//var ok1, ok2 bool
	for {
		select {
		// 5-7s之内，从一个已关闭的 channel 接收数据将永远不会被阻塞
		case x, ok := <-ch1:
			if !ok {
				//判断 ch1 或 ch2 被关闭后，显式地将 ch1 或 ch2 置为 nil。
				ch1 = nil
			} else {
				fmt.Println(x)
			}
			//fmt.Println(x) // 5-7s 一直输出0
		case x, ok := <-ch2:
			if !ok {
				//判断 ch1 或 ch2 被关闭后，显式地将 ch1 或 ch2 置为 nil。
				ch2 = nil
			} else {
				fmt.Println(x)
			}
		}

		//if ok1 && ok2 {
		//	break
		//}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")
}

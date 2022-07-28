package foo

import (
	"testing"
)

// for msend benchmark test
var mc1 chan string

// for mrecv benchmark test
var mc2 chan string

func init() {
	mc1 = make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				<-mc1
			}
		}()
		go func() {
			for {
				mc1 <- "hello"
			}
		}()
	}

	mc2 = make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				mc2 <- "hello"
			}
		}()
		go func() {
			for {
				<-mc2
			}
		}()
	}
}

func msend(msg string) {
	mc1 <- msg
}
func mrecv() {
	<-mc2
}

func BenchmarkUnbufferedChanNToNmsend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		msend("hello")
	}
}
func BenchmarkUnbufferedChanNToNmrecv(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mrecv()
	}
}

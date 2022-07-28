package main

import "fmt"

type T int

func (t T) Error() string {
	return "bad error"
}

func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	println("err1:", err1)

	err1 = (*T)(nil)
	println("err1:", err1)
	println("err1 = nil:", err1 == nil)

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
}

func main() {
	printNonEmptyInterface()
	//err1: (0x1043481b8,0x0)
	//	err1 = nil: false
	//err1: (0x104348218,0x10432f178)
	//err2: (0x104348218,0x10432f180)
	//	err1 = err2: false
	//err1: (0x104348218,0x10432f178)
	//err2: (0x104348138,0x14000010230)
	//	err1 = err2: false

}

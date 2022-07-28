package main

import (
	"GoFirstClass/tool"
)

type T struct{}

func (T) M1()  {}
func (T) M2()  {}
func (*T) M3() {}
func (*T) M4() {}

type Interface interface {
	M1()
	M2()
}

func main() {
	//var n int
	//dumpMethodSet(n)
	//dumpMethodSet(&n)
	//var t T
	//dumpMethodSet(t)
	//dumpMethodSet(&t) // *T 类型的方法集合包含所有以 *T 为 receiver 参数类型的方法，以及所有以 T 为 receiver 参数类型的方法

	type S T
	var s1 S
	tool.DumpMethodSet(s1)  // main.S's method set is empty!
	tool.DumpMethodSet(&s1) // *main.S's method set is empty!

}

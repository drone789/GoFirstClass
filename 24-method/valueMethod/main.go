package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

//func (p *field) print() {
//	fmt.Println(p.name)
//}

func (p field) print() {
	fmt.Println(p.name)
}
func main() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		//go v.print() // one two  three

		// 通过方法表达式，以上等价于
		// 切片中元素类型是*field指针,v就是元素的地址
		// 与 print 的 receiver 参数类型相同，
		// 每次调用 (*field).print 函数时直接传入的 v 即可，实际上传入的也是各个 field 元素的地址
		//go (*field).print(v)
		go field.print(*v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		//go v.print()

		go field.print(v)
	}

	time.Sleep(3 * time.Second)
}

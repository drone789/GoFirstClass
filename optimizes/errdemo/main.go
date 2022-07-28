package main

import (
	"GoFirstClass/optimizes/errdemo/sub1"
	"fmt"
)

func main() {
	err := sub1.Diff(1, 2)
	//fmt.Println(err)
	// 在pkg/error中打印堆栈信息
	fmt.Printf("%+v", err)
}

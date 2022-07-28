package main

func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
}

func main() {
	printNilInterface()
	//(0x0,0x0)
	//(0x0,0x0)
	//i = nil: true
	//err = nil: true
	//i = err: true

}

package main

import "unsafe"

type T int

func (t T) Error() string {
	return "bad error"
}

func printEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)
}

func main() {
	//printEmptyInterfaceAndNonEmptyInterface()

	//eif: (0x1029710c0,0x102968cf0)
	//err: (0x102977f18,0x102968cf0)

	// Go 在进行等值比较时，类型比较使用的是 eface 的 _type 和 iface 的 tab._type，
	//	因此就像我们在这个例子中看到的那样，当 eif 和 err 都被赋值为T(5)时，两者之间是划等号的。
	//	eif = err: true

	//eif: (0x1029710c0,0x102968cf0)
	//err: (0x102977f18,0x102968cf8)
	//	eif = err: false

	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err)
	//dumpEface(eif)
	dumpItabOfIface(unsafe.Pointer(&err))
	dumpDataOfIface(err)

}

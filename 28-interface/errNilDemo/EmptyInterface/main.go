package main

func printEmptyInterface() {
	var eif1 interface{} // 空接口类型
	var eif2 interface{} // 空接口类型
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // false

	eif2 = 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // true

	eif2 = int64(17)
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 = eif2:", eif1 == eif2) // false
}

func main() {
	printEmptyInterface()
	//eif1: (0x100a47e00,0x14000098f60)
	//eif2: (0x100a47e00,0x14000098f58)
	//	eif1 = eif2: false
	//eif1: (0x100a47e00,0x14000098f60)
	//eif2: (0x100a47e00,0x100a40be8)
	//	eif1 = eif2: true
	//eif1: (0x100a47e00,0x14000098f60)
	//eif2: (0x100a47ec0,0x100a40be8)
	//	eif1 = eif2: false

}

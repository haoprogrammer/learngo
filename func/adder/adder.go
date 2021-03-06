package main

import "fmt"

//累加器
func adder() func(i int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

//简单纯函数编程
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	//
	//a := adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(a(i))
	//}

	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}

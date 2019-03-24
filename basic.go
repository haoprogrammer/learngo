package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue(){
	var a int
	var s string
	fmt.Print(a,s)
}

func variableIntiValue(){

	var a, b int = 3, 5
	var s string = "hello"
	fmt.Printf("\n%d %d %q",a, b, s)
}


func euler(){
	c := 3 +4i
	fmt.Println(c)
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)

	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)

}

func main() {
	variableZeroValue()
	variableIntiValue()

	euler()
}

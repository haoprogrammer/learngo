package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {

	//i, err := strconv.Atoi("-42")
	fmt.Println(convertToBin(5))
	fmt.Println(s)
}

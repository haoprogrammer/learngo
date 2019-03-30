package main

import (
	"fmt"
)

func eval(a, b int, op string) (int, error) {

	switch op {
	case "+":
		return a + b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}

}

func sum(numbers ...int) int {
	result := 0
	for i := range numbers {
		result += numbers[i]
	}
	return result

}

func swap(a, b int) (int, int) {
	return b, a
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int), a, b int) {

}

func main() {
	if result, err := eval(3, 4, "+"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

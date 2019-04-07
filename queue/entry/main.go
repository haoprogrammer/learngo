package main

import (
	"fmt"
	"haoprogrammer/learngo/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

}

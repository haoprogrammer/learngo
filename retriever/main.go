package main

import (
	"fmt"
	"haoprogrammer/learngo/retriever/mock"
	"haoprogrammer/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {

	return r.Get("http://www.imooc.com")
}

func main() {

	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)

	//fmt.Println(download(r))
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

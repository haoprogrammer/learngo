package main

import (
	"fmt"
	"time"
)

/**
goroutine 可能切换的点
1. I/O，select
2. 函数调用
3. channel
4. 等待锁
5. runtime.Gosched()
作为参考，不能保证切换，不能保证其他地方不切换
*/
func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println("hello from goroutine"+"goroutine %d\n", i)
				//runtime.Gosched()
			}
		}(i)
	}
	//
	time.Sleep(time.Minute)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2 ; i++  {
			fmt.Println("子协程 i = ",i)
			time.Sleep(time.Second)
		}
		ch <- "123"
	}()

	str := <-ch //没有数据会堵塞
	fmt.Println("str = " ,str)
}
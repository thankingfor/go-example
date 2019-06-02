package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("ccccccccccccc")

	runtime.Goexit()//终止所在的协程
	fmt.Println("dddddddddddddddddd")
}

func main()  {

	//创建新建的协程
	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaaaaa")

		test()

		fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	}()

	//特别写一个死循环 目的不让主线程结束
	for {
	}

}

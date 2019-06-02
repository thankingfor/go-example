package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go func() {
		for i:=1;i<5 ;i++  {
			fmt.Print("go")
		}
	}()

	for i:=1;i<2 ;i++  {
		//让出时间片 先让别的协程执行，它执行玩，再回来执行协程
		runtime.Gosched()
		fmt.Print("hello")
	}
}
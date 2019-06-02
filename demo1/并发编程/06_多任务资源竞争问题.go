package main

import (
	"fmt"
	"time"
)
//全局变量  创建一个channel
var ch = make(chan int)

func Printer(str string)  {
	for _,data := range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func main()  {

	go func() {
		Printer("hello")
		ch <- 666// 给管道写数据 发送
	}()

	go func() {
		<-ch //从管道去数据，接受 如果通道没数据就不接受
		Printer("world")
	}()

	for{}
}

package main

import "fmt"


func main() {
	ch := make(chan int)//创建一个无缓存channel

	go func() {
		for i := 0; i < 5 ; i++  {
			ch <- i //往通道写数据
		}
		//不需要写数据时
		close(ch)
	}()

	for num := range ch{
		fmt.Println("num = ",num);
	}
}

func main01() {
	ch := make(chan int)//创建一个无缓存channel

	go func() {
		for i := 0; i < 5 ; i++  {
			ch <- i //往通道写数据
		}
		//不需要写数据时
		close(ch)
	}()

	for {
		//如果为ture 管道没有关闭
		if num , ok :=  <- ch; ok == true{
			fmt.Println("num = ",num);
		} else {
			break
		}
	}
}

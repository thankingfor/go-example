package main

import "fmt"

//此通道只能写 不能读
func producer(out chan <- int)  {
	for i := 0;i<10 ;i++  {
		out <- i*i
	}
	close(out)
}
//只能读 不能写
func consumer(in <-chan int)  {
	for num := range in{
		fmt.Println(num)
	}
}

func main() {
	//创建一个channel 双向的
	ch := make(chan int)

	//生产者 生产数字 写入channel
	go producer(ch)

	//消费者 从channel读取内容 打印
	consumer(ch)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器 设置时间为2s 2s后  往time通道里写内容（当前时间）
	timer := time.NewTimer(2 * time.Second)  //NewTimer只会响应一次
	fmt.Println("当前时间 ：",time.Now())
	t := <- timer.C //2s后 往timer.C写数据  有数据后 就可以读取  channel没有数据前后堵塞
	fmt.Println("t = ",t)
}

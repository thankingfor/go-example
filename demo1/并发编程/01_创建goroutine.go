package main

import (
	"fmt"
	"time"
)

func newTask()  {
	for  {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}

func main()  {

	go newTask() //新建以恶搞写成 新建一个任务

	for  {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}

}

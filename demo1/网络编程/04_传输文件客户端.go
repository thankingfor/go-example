package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("请输入需要传输的文件: ")
	var  path string
	fmt.Scan(&path)

	//获取文件名
	info ,err := os.Stat(path)
	PrintErr(err)

	//主动链接服务器
	conn , err1 := net.Dial("tcp","127.0.0.1:8000")
	PrintErr(err1)
	defer conn.Close()

	//给接收方发送文件名
	_,err2 :=conn.Write([]byte(info.Name()))
	PrintErr(err2)

	//接受回复


}

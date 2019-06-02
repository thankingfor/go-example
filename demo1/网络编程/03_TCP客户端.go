package main

import (
	"fmt"
	"net"
	"os"
)




func main() {
	//主动链接服务器
	conn ,_ := net.Dial("tcp","127.0.0.1:8000")
	defer conn.Close()
	go func() {
		//切片缓存
		buf := make([]byte,1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("err = ",err)
				return
			}
			fmt.Println(string(buf[:n]))//打印接收到的数据
		}
	}()

	str := make([]byte,1024)
	for  {
		n, _ := os.Stdin.Read(str)//从键盘读取内容
		//发送数据
		conn.Write(str[:n])
	}

}

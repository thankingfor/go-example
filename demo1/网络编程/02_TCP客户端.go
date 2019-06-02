package main

import "net"

func main() {
	//主动链接服务器
	conn ,_ := net.Dial("tcp","127.0.0.1:8000")
	defer conn.Close()

	//发送数据
	conn.Write([]byte("are u ok?"))
}

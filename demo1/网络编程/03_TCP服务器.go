package main

import (
	"fmt"
	"net"
	"strings"
)

//处理用户请求
func HandleConn(conn net.Conn)  {
	//函数调用完毕 自动关闭conn
	defer conn.Close()

	//获取客户端的网络地址信息
	addr :=  conn.RemoteAddr().String()
	fmt.Println("addr connect successful"+addr)

	//读取用户数据
	buf := make([]byte,2024)
	n,err := conn.Read(buf)
	if err != nil {
		fmt.Println("err = ",err)
		return
	}
	fmt.Println("read buf = ",string(buf[:n]))
	if "exit" == string(buf[:n]){
		fmt.Println(addr , "exit!")
		return
	}

	//把数据转为大写
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

}

func main() {
	//监听
	listener,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户连接
	for {
		conn , err := listener.Accept()
		if err != nil {
			fmt.Println("err = ",err)
			return
		}
		//接受用户的请求
		go  HandleConn(conn)
	}

}

package main

import (
	"fmt"
	"net"
)

type Client struct{
	C chan string //用户发送数据的通道
	Name string //用户名
	Addr string //网络地址
}

//保存在线用户
var onlineMap map[string]Client
var message = make(chan string)

func MakeMsg(cli Client,msg string)(buf string)  {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return
}

func HandleConns(conn net.Conn)  {
	defer conn.Close()
	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()
	//创建一个结构体
	cli := Client{make(chan string),cliAddr,cliAddr}
	onlineMap[cliAddr] = cli

	//新开一个协程 专门给当前客户端发送消息
	go WriteMsgToClient(cli,conn)

	//广播某个人在线
	message <- MakeMsg(cli,"login")

	go func() {
		buf := make([]byte,2048)
		for {
			n,err := conn.Read(buf)
			if n == 0 { //对方端口 或者粗问题
				fmt.Println("conn Read err = ",err)
				return
			}
			msg := string(buf[:n-1])
			//转发此内容
			message <- MakeMsg(cli,msg)
		}
	}()



	for   {

	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C{//给当前客户端发送消息
		conn.Write([]byte(msg + "\n"))
	}
}

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <- message  //没有消息钱 这里会阻塞
		//遍历map 给每个map成员都发送次消息
		for _,cli := range onlineMap{
			cli.C <- msg
		}
	}
}


func main() {
	listener, e := net.Listen("tcp", ":8000")
	if e != nil {
		fmt.Println("listener.Accept err = ",e)
		return
	}
	defer listener.Close()

	//新开一个协程转发消息 只要有消息来了 变量map 给map每个成员都发送此消息
	go Manager();

	//主协程用于接受用户来链接
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ",err)
			continue
		}
		go HandleConns(conn)
	}
}


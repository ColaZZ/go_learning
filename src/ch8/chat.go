package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster() // 广播器
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handlerConn(conn) // 连接处理
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有接收的客户消息
)

func broadcaster() {
	clients := make(map[client]bool) //所有连接的客户端
	for {
		select {
		case msg := <-messages:
			//把所有接收的消息广播给所有的客户
			//发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handlerConn(conn net.Conn) {
	ch := make(chan string) // 对外发送客户信息的通道
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + "has arrived"
	entering <- ch // 通过entering通道通知广播者新客户到来

	input := bufio.NewScanner(conn) //读取客户发来的每一行文本
	for input.Scan() {
		messages <- who + ":" + input.Text() //通过全局接收消息通道将每一行发送给广播者 ID+消息
	}

	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

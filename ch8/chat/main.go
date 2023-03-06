package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {

	server, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
		return
	}
	go broadcast()

	for {
		// main goroutine : 处理客户端连接请求
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcast goroutine : 处理消息分发给各个客户端，管理客户端的新增记录和减除回收
func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
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

// 同时处理 来自于客户端的消息 和 发回给客户端的消息
// 对于一个新增的连接，初始化客户端的管道资源，并发送回[broadcast goroutine],进行集中处理
// 同时，启动一个对应的[back-write goroutine]用于给客户端回写消息
// 最后，接收客户端管道发来的消息，进行显示处理
// 如果客户端资源关闭，则通知[broadcast goroutine]进行资源回收
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(ch, conn)

	who := conn.RemoteAddr().String()
	ch <- "欢迎您来到我们的聊天室 : " + who
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + " :" + input.Text()
	}

	leaving <- ch
	messages <- who + " 已经离开了"
	conn.Close()
}

func clientWriter(ch <-chan string, conn net.Conn) {
	for msg := range ch {
		//io.Copy(conn, strings.NewReader(msg))
		fmt.Fprintln(conn, msg)
	}
}

package ch8

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type ClockServer int64

func (cs ClockServer) Run() {
	if cs == 0 {
		clock1Server()
	} else if cs == 1 {
		clock2Server()
	} else if cs == 2 {
		revert()
	}
}

var clients int

func clock1Server() {
	server, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := server.Accept()
		clients++
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		fmt.Printf("\tserver : 连接 - %d即将关闭\n", clients)
		conn.Close()
	}()
	fmt.Printf("\tserver : 连接 - %d 开始输出\n", clients)
	for {
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			fmt.Errorf("conn was disabled for : %v", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func clock2Server() {
	server, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

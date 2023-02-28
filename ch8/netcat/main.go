package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	//netcat1(conn)
	netcat3(conn)
}

func netcat1(conn net.Conn) {
	defer conn.Close()
	go mustCopyTo(os.Stdout, conn)
	mustCopyTo(conn, os.Stdin)
}

func mustCopyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalf("client output error : %v", err)
	}
}

/// netcat version 3
func netcat3(conn net.Conn) {
	ch := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done!")
		ch <- struct{}{}
	}()
	mustCopyTo(conn, os.Stdin)
	conn.Close()
	<-ch
}

package ch8

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func revert() {
	log.Println("launch revert version 1 server...")
	server, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn2(conn)
	}
}

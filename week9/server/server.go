package main

import (
	"fmt"
	"gowork/week9/protocol"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		p := protocol.Proto{}
		err := p.Read(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%t\n", p)
		fmt.Println(string(p.Body))
	}
}

func main() {
	l, err := net.Listen("tcp", ":2300")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		fmt.Println("建立连接")
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

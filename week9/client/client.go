package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gowork/week9/protocol"
)

func main() {
	conn, err := net.Dial("tcp", ":2300")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		input := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(input)
			if err != nil {
				fmt.Println("input err:", err)
				continue
			}
			p := protocol.Proto{
				Op:   protocol.OpAuth,
				Body: input[:n],
			}
			p.Write(conn)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleData(conn)
	}
}

func handleData(c net.Conn) {
	defer c.Close()

	var b [512]byte
	for {
		n, err := c.Read(b[:])
		if err != nil {
			break
		}
		fmt.Println(b[:n])
	}
}

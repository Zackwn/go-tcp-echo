package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go Echo(conn)
	}
}

func Echo(conn net.Conn) {
	defer conn.Close()
	for {
		echo := make([]byte, 1024)
		n, err := conn.Read(echo)
		if err != nil {
			log.Println(err)
			break
		}

		_, err = conn.Write(echo[:n])
		if err != nil {
			log.Println(err)
			break
		}
	}
}

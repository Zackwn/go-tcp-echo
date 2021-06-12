package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		var str string
		fmt.Scanln(&str)
		// sending nothing will result in a bug
		// where both server and client will be waiting to read
		if len(str) == 0 {
			str = " "
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			log.Println(err)
			break
		}

		buffer := make([]byte, 1024)
		bufferSize, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(string(buffer[:bufferSize]))
	}
}

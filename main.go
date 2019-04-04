package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	PORT := ":" + "8001"

	var store sync.Map

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepting new connection!")
			panic(err)
		}

		fmt.Println("New connection")
		go (&sessionHandler{conn, &store}).handle()
	}
}

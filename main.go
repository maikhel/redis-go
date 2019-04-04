package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/maikhel/redis-go/redis"
)

var store sync.Map

func main() {
	PORT := ":" + "8001"

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepting new connection!")
			panic(err)
		}

		fmt.Println("New connection")
		go redis.HandleConnection(connection, &store)
	}
}

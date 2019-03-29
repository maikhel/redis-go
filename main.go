package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/maikhel/redis-go/redis"
)

var store sync.Map

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(1)
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error while accepting new connection!")
			// panic(err)
		}
		go redis.HandleConnection(c, &store)
	}
}

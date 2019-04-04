package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type specification struct {
	Port            int    `envconfig:"PORT" default:"6379"`
	DefaultPassword string `envconfig:"REDIS_AUTH_PASS" default:"bacon"`
}

var cfg specification

func init() {
	envconfig.MustProcess("", &cfg)
}

func main() {
	var store sync.Map

	port := fmt.Sprintf(":%d", cfg.Port)
	listener, err := net.Listen("tcp", port)
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
		go (&sessionHandler{conn, &store, false}).handle()
	}
}

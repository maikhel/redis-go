package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

var store map[string]string

func ping(input string) string {
	if input != "" {
		return input
	} else {
		return "PONG"
	}
}

func execCommand(input string) string {
	args := strings.Split(input, " ")
	cmd := args[0]

	switch cmd {
	case "ping":
		{
			var arg string
			if len(input) > 5 {
				arg = input[5:]
			} else {
				arg = ""
			}
			return ping(arg)
		}
	case "get":
		{
			if len(args) < 2 {
				return "-ERR Too few arguments"
			} else {
				return store[args[1]]
			}
		}
	case "set":
		{
			if len(args) < 3 {
				return "-ERR Too few arguments"
			} else {
				key, val := args[1], args[2]
				store[key] = val
				return "+OK"
			}
		}
	default:
		return "COMMAND NOT FOUND"
	}

}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		input := strings.TrimSpace(string(netData))
		var result string
		if input == "STOP" {
			break
		} else {
			result = execCommand(input) + "\n"
		}

		c.Write([]byte(string(result)))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	store = make(map[string]string)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

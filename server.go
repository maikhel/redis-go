package main

import (
	"bufio"
	"fmt"
	"io"
	"net/textproto"
	"strings"
	"sync"
)

type sessionHandler struct {
	conn  io.ReadWriteCloser
	store *sync.Map
	auth  bool
}

func (s *sessionHandler) handle() {
	defer fmt.Println("Connection closed")
	defer s.conn.Close()

	buffer := bufio.NewReader(s.conn)

	for {
		netData, err := textproto.NewReader(buffer).ReadLine()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}

		input := strings.TrimSpace(string(netData))

		var result string
		result = s.ExecCommand(input, s.store) + "\n"

		s.conn.Write([]byte(string(result)))
	}
}

func (s *sessionHandler) ExecCommand(input string, store *sync.Map) string {
	args := strings.Split(input, " ")
	cmd := args[0]

	if !s.auth && cmd != "auth" {
		return "-NOAUTH Authentication required."
	}

	switch cmd {
	case "auth":
		{
			if len(args) < 2 {
				return "-ERR Too few arguments"
			}
			if s.auth {
				return "-ERR alrady authenticated."
			}

			if args[1] == cfg.DefaultPassword {
				s.auth = true
				return "+OK"
			}
			return "-ERR invalid password"
		}
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
			}
			result, ok := store.Load(args[1])
			if ok {
				return result.(string)
			}
			return "(nil)"
		}
	case "set":
		{
			if len(args) < 3 {
				return "-ERR Too few arguments"
			}
			store.Store(args[1], args[2])
			return "+OK"
		}
	default:
		return "COMMAND NOT FOUND"
	}

}

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

		result = ExecCommand(input, s.store) + "\n"

		s.conn.Write([]byte(string(result)))
	}
}

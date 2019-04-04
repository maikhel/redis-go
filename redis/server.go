package redis

import (
	"bufio"
	"fmt"
	"io"
	"net/textproto"
	"strings"
	"sync"
)

func HandleConnection(c io.ReadWriteCloser, store *sync.Map) {
	defer fmt.Println("Closing connection")
	defer c.Close()

	buffer := bufio.NewReader(c)

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

		result = ExecCommand(input, store) + "\n"

		c.Write([]byte(string(result)))
	}
}

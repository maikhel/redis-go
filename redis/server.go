package redis

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func HandleConnection(c net.Conn, store *sync.Map) {

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
		}

		result = ExecCommand(input, store) + "\n"

		c.Write([]byte(string(result)))
	}
	defer c.Close()
}

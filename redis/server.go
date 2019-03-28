package redis

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConnection(c net.Conn) {
	var store map[string]string
	store = make(map[string]string)

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

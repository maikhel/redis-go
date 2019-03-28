package redis

import (
	"strings"
)

func ping(input string) string {
	if input != "" {
		return input
	}
	return "PONG"
}

func ExecCommand(input string, store map[string]string) string {
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
			}
			return store[args[1]]
		}
	case "set":
		{
			if len(args) < 3 {
				return "-ERR Too few arguments"
			}
			key, val := args[1], args[2]
			store[key] = val
			return "+OK"
		}
	default:
		return "COMMAND NOT FOUND"
	}

}

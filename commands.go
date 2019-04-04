package main

import (
	"strings"
	"sync"
)

func ping(input string) string {
	if input != "" {
		return input
	}
	return "PONG"
}

func ExecCommand(input string, store *sync.Map) string {
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

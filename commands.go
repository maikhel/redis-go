package main

func ping(input string) string {
	if input != "" {
		return input
	}
	return "PONG"
}

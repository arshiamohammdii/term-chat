package main

import (
	"fmt"
	"log"

	"term-chat.com/server"
)

func main() {
	const esc = "\x1b"

	const (
		Reset     = esc + "[0m"
		Bold      = esc + "[1m"
		Dim       = esc + "[2m"
		Italic    = esc + "[3m"
		Underline = esc + "[4m"
		Blink     = esc + "[5m"
		Invert    = esc + "[7m"
	)
	fmt.Print(Bold + "hello" + Reset)
	server := server.NewServer(":2222", "/Users/arshiamohammadi/.ssh/host_key")
	log.Fatal(
		server.ListenSSH(),
	)
}

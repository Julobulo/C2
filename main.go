package main

import (
	"github.com/Julobulo/C2/c2"
)

func main() {
	listener := c2.HttpListener{
		Ip:   "127.0.0.1",
		Port: "1234",
	}

	// Run listener as go routine
	go listener.StartListener()

	// Start CLI
	c2.StartCLI()
}

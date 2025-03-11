package main

import (
	"fmt"
	"os"
	"tcp-congestion-control/client"
	"tcp-congestion-control/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|client]")
		return
	}

	switch os.Args[1] {
	case "server":
		server.StartServer()
	case "client":
		client.StartClient()
	default:
		fmt.Println("Invalid argument. Use 'server' or 'client'.")
	}
}



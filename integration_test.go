package main

import (
	"bufio"
	"net"
	"testing"
	"time"
	"tcp-congestion-control/client"
	"tcp-congestion-control/server"
)

func TestClientServerIntegration(t *testing.T) {
	// Start the server in a separate goroutine
	go server.StartServer()
	time.Sleep(1 * time.Second) // Give the server some time to start

	// Start the client
	go client.StartClient()
	time.Sleep(1 * time.Second) // Give the client some time to start

	// Connect to the server as a mock client
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Send a test message
	conn.Write([]byte("test message\n"))

	// Read the response
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}

	expectedResponse := "Echo: test message\n"
	if response != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, response)
	}
}


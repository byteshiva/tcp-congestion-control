package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
	"tcp-congestion-control/congestion"
)

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	cc := congestion.NewCongestionController(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		err := cc.SendData([]byte(line + "\n"))
		if err != nil {
			fmt.Println("Error sending data:", err)
			cc.HandleTimeout()
			continue
		}
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			cc.HandleTimeout()
			continue
		}
		fmt.Println("Received:", response)
		cc.HandleAck()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}


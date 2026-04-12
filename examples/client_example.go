package examples

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// Example of how to use the new JSON protocol
func example() {
	// Connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:6767")
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	fmt.Println("Connected to Input Engine")

	// Example 1: Press 'A' key for 50ms
	sendInput(conn, "VK_A", 50)

	time.Sleep(100 * time.Millisecond)

	// Example 2: Left click for 10ms
	sendInput(conn, "VK_LBUTTON", 10)

	time.Sleep(100 * time.Millisecond)

	// Example 3: Press Enter key for 100ms
	sendInput(conn, "VK_RETURN", 100)

	fmt.Println("Done!")
}

func sendInput(conn net.Conn, virtualKey string, duration uint16) {
	request := map[string]interface{}{
		"virtual_key": virtualKey,
		"duration":    duration,
	}

	jsonData, _ := json.Marshal(request)
	fmt.Printf("Sending: %s\n", string(jsonData))

	conn.Write(append(jsonData, '\n'))
}

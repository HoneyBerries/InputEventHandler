package listener

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"InputEventHandler/input"
	"InputEventHandler/keymap"
)

// InputRequest represents a JSON input request
type InputRequest struct {
	VirtualKey string `json:"virtual_key"`
	Duration   uint16 `json:"duration"`
}

// Start begins listening for connections on the specified port
func Start(port int) error {
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return err
	}
	fmt.Println(">>> Input Engine Active <<<")
	fmt.Printf("Listening on 127.0.0.1:%d\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

// handleConnection processes a single client connection
func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)
	fmt.Printf("Connected: %s\n", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()

		var req InputRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			fmt.Printf("Failed to parse JSON: %v\n", err)
			continue
		}

		keyDef, ok := keymap.VirtualKeyMap[req.VirtualKey]
		if !ok {
			fmt.Printf("Unknown virtual key: %s\n", req.VirtualKey)
			continue
		}

		fmt.Printf("VirtualKey=%s Duration=%dms\n", req.VirtualKey, req.Duration)
		input.Tap(keyDef, req.Duration)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Disconnected:", err)
	}
}

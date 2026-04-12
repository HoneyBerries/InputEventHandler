package main

import (
	"fmt"

	"InputEventHandler/listener"
)

func main() {
	port := 6767
	if err := listener.Start(port); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

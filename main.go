package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"InputEventHandler/listener"
)

func main() {
	// Parse command line flags
	port := flag.Int("port", 6767, "port to listen on")
	flag.Parse()

	// Configure logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Validate port
	if *port < 1 || *port > 65535 {
		log.Fatalf("Invalid port: %d (must be 1-65535)", *port)
	}

	// Create server
	srv := listener.NewServer(*port)

	// Set up graceful shutdown on interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	sig := <-sigChan
	log.Printf("Received signal: %v", sig)
	log.Print("Shutting down...")

	srv.Shutdown()
}

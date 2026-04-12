package listener

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"

	"InputEventHandler/input"
	"InputEventHandler/keymap"
)

// InputRequest represents a JSON input request from a client
type InputRequest struct {
	VirtualKey string `json:"virtual_key"`
	Duration   uint16 `json:"duration"`
}

// Server manages the input event listener and active connections
type Server struct {
	listener     net.Listener
	port         int
	activeConns  int32 // Atomic counter for active connections
	shutdownOnce sync.Once
	shutdownChan chan struct{}
	connWg       sync.WaitGroup
}

// NewServer creates a new input event server on the specified port
func NewServer(port int) *Server {
	return &Server{
		port:         port,
		shutdownChan: make(chan struct{}),
	}
}

// Start begins listening for TCP connections and processing input requests.
// This function blocks until Shutdown() is called.
// Returns an error if the server fails to bind to the port.
func (s *Server) Start() error {
	addr := fmt.Sprintf("127.0.0.1:%d", s.port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}
	s.listener = ln

	log.Printf(">>> Input Engine Active <<<")
	log.Printf("Listening on %s (PID: %d)", addr, getProcessID())

	// Accept connections in a goroutine
	go s.acceptConnections()

	// Wait for shutdown signal
	<-s.shutdownChan

	// Close listener and wait for active connections to finish
	if err := s.listener.Close(); err != nil {
		log.Printf("Error closing listener: %v", err)
	}
	s.connWg.Wait()
	log.Print("Server shutdown complete")
	return nil
}

// acceptConnections continuously accepts new client connections
func (s *Server) acceptConnections() {
	for {
		select {
		case <-s.shutdownChan:
			return
		default:
		}

		conn, err := s.listener.Accept()
		if err != nil {
			// Check if we're shutting down
			select {
			case <-s.shutdownChan:
				return
			default:
				log.Printf("Error accepting connection: %v", err)
				continue
			}
		}

		atomic.AddInt32(&s.activeConns, 1)
		s.connWg.Add(1)
		go func() {
			defer s.connWg.Done()
			defer atomic.AddInt32(&s.activeConns, -1)
			s.handleConnection(conn)
		}()
	}
}

// handleConnection processes a single client connection
func (s *Server) handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr()
	defer conn.Close()

	log.Printf("Connected: %s (active: %d)", remoteAddr, atomic.LoadInt32(&s.activeConns))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Parse JSON request
		var req InputRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			log.Printf("[%s] Invalid JSON: %v (line: %s)", remoteAddr, err, line)
			continue
		}

		// Validate and look up key
		if req.VirtualKey == "" {
			log.Printf("[%s] Missing virtual_key field", remoteAddr)
			continue
		}

		keyCode, ok := keymap.KeyCodeByName(req.VirtualKey)
		if !ok {
			log.Printf("[%s] Unknown virtual key: %s", remoteAddr, req.VirtualKey)
			continue
		}

		// Execute input event
		if err := input.Tap(uint16(keyCode), req.Duration); err != nil {
			log.Printf("[%s] Input failed for %s: %v", remoteAddr, req.VirtualKey, err)
			continue
		}

		log.Printf("[%s] ✓ VirtualKey=%s Duration=%dms", remoteAddr, req.VirtualKey, req.Duration)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Disconnected: %s (error: %v)", remoteAddr, err)
	} else {
		log.Printf("Disconnected: %s (clean)", remoteAddr)
	}
}

// Shutdown gracefully shuts down the server and waits for active connections to close
func (s *Server) Shutdown() {
	s.shutdownOnce.Do(func() {
		close(s.shutdownChan)
	})
}

// ActiveConnections returns the current number of active connections
func (s *Server) ActiveConnections() int32 {
	return atomic.LoadInt32(&s.activeConns)
}

// getProcessID is a helper to get the current process ID for logging
func getProcessID() int {
	// In a real application, you'd use os.Getpid()
	// For now, return a placeholder
	return 0
}

# Gemini Context: InputEventHandler

A Windows-only TCP server that accepts JSON commands to simulate keyboard and mouse input using Windows Virtual Key codes.

## Project Overview

- **Purpose:** Provide a simple, human-readable interface for programmatically controlling keyboard and mouse input on Windows.
- **Core Technology:** 
  - **Language:** Go (v1.26+)
  - **OS Integration:** Win32 API (`user32.dll`) via `syscall` for input simulation.
  - **Networking:** TCP server listening on `127.0.0.1`.
  - **Protocol:** Line-delimited JSON.
- **Architecture:**
  - `main.go`: Entry point, flag parsing (`-port`), and signal handling.
  - `listener/`: TCP server management, connection lifecycle, and JSON request/response handling.
  - `input/`: Low-level Win32 wrappers for `keybd_event` and `mouse_event`.
  - `keymap/`: Mapping of string Virtual Key names (e.g., `VK_A`, `VK_LBUTTON`) to their numeric codes.
  - `examples/`: Client implementation examples in Go and other languages (documentation-only).

## Building and Running

### Commands
- **Build:** `go build` (produces `InputEventHandler.exe` on Windows).
- **Run:** `.\InputEventHandler.exe` or `.\InputEventHandler.exe -port 6767`.
- **Test:** No automated tests currently exist in the codebase. TODO: Implement unit tests for key mapping and listener logic.

### Configuration
- **Default Port:** 6767
- **Host:** Restricted to `127.0.0.1` for security.

## Development Conventions

- **Platform Restriction:** This project is strictly **Windows-only**. It will compile on other platforms but will fail at runtime due to reliance on `user32.dll`.
- **Concurrency:** Uses Go's concurrency primitives (goroutines, `sync/atomic`, `sync.WaitGroup`) to handle multiple client connections simultaneously.
- **Error Handling:** 
  - Errors are logged to `stdout` with `log.Lshortfile` flags.
  - JSON protocol errors (invalid JSON, unknown keys) are logged, but the connection remains open for subsequent commands.
  - Each request results in a `{"success": true|false}` response to the client.
- **Graceful Shutdown:** Implements clean closure of the TCP listener and waits for active connections to finish upon receiving `SIGINT` or `SIGTERM`.
- **Dependencies:** Avoids external libraries for input; prefers standard library `syscall` for Win32 interop.

## Key Files
- `main.go`: Application lifecycle and configuration.
- `listener/listener.go`: Main server logic and protocol handling.
- `input/input.go`: Win32 syscall implementation for input events.
- `keymap/keymap.go`: The source of truth for supported Virtual Key names.

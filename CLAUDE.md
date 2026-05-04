# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build -o InputEventHandler.exe

# Run (default port 6767)
.\InputEventHandler.exe

# Run on a custom port
.\InputEventHandler.exe -port 8080

# Test manually with netcat
echo '{"virtual_key":"VK_A","duration":50}' | nc localhost 6767
```

There are no automated tests. The CI workflow (`.github/workflows/build.yml`) only builds the binary on `windows-latest`.

## Architecture

Four packages with a clear single-direction dependency chain:

```
main → listener → input
              └→ keymap
```

- **`main.go`** — parses the `-port` flag, wires up `SIGINT`/`SIGTERM` for graceful shutdown, and delegates everything to `listener.Server`
- **`listener/listener.go`** — TCP server; accepts concurrent connections (one goroutine per client), reads line-delimited JSON, validates and looks up the key via `keymap`, calls `input.Tap`, and writes a `{"success": bool}` response per request
- **`input/input.go`** — thin Win32 wrapper around `keybd_event` and `mouse_event` from `user32.dll` via Go's `syscall` package; the only place that touches the OS
- **`keymap/keymap.go`** — static map of `VK_*` name strings → numeric codes; keyboard uses standard Windows virtual key codes (0x01–0xFF), mouse uses custom codes 256–260 to avoid overlap

## Key Design Details

**Mouse vs keyboard code space:** Mouse button codes start at 256 to stay safely above the 0xFF ceiling of Windows virtual key codes. `input.Tap` dispatches to `tapMouse` vs `tapKeyboard` based on whether `keyCode >= 256`.

**Windows success-as-error quirk:** `syscall` on Windows returns a non-nil error even on success with the message `"The operation completed successfully."` — `isSuccessError` in `input/input.go` detects this and suppresses it.

**Protocol ordering:** The server sends one JSON response per request line. Clients must read the response before sending the next request; this is the only backpressure mechanism and ensures the previous input has completed before the next one starts.

**Windows-only:** The binary will only compile and run on Windows. There are no build tags enforcing this — it will simply fail to link on other platforms due to the `user32.dll` imports.

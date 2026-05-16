# InputEventHandler

[![Build](https://github.com/HoneyBerries/InputEventHandler/actions/workflows/build.yml/badge.svg)](https://github.com/HoneyBerries/InputEventHandler/actions/workflows/build.yml)
[![Latest Release](https://img.shields.io/github/v/release/HoneyBerries/InputEventHandler?sort=semver)](https://github.com/HoneyBerries/InputEventHandler/releases/latest)
[![Release Workflow](https://github.com/HoneyBerries/InputEventHandler/actions/workflows/release.yml/badge.svg)](https://github.com/HoneyBerries/InputEventHandler/actions/workflows/release.yml)
[![Platform](https://img.shields.io/badge/platform-Windows-blue)](#)

Windows-only TCP server that accepts line-delimited JSON commands and simulates keyboard/mouse input using Windows Virtual Key (`VK_*`) names.

## Quick start

1. Download `InputEventHandler.exe` from the latest GitHub release.
2. Run it (default port is `6767`):

   ```powershell
   .\InputEventHandler.exe
   ```

3. Send a command (one JSON object per line) and read the response:

   ```bash
   echo '{"virtual_key":"VK_A","duration":50}' | nc 127.0.0.1 6767
   ```

## Protocol

Each request line is a JSON object:

```json
{"virtual_key":"VK_A","duration":50}
```

- `virtual_key` (string, required): key name like `VK_A`, `VK_RETURN`, `VK_LBUTTON`
- `duration` (uint16, required): hold duration in milliseconds (use `0` for “tap now”)

Each request gets exactly one JSON response line:

```json
{"success":true}
```

Clients should read the response before sending the next request (this is the server’s backpressure mechanism).

## Virtual keys

- Keyboard keys use standard Windows virtual-key codes (`VK_A`, `VK_F1`, `VK_RETURN`, …).
- Mouse buttons are exposed as `VK_LBUTTON`, `VK_RBUTTON`, `VK_MBUTTON`, `VK_XBUTTON1`, `VK_XBUTTON2`.
- Full supported key list lives in `keymap.go`.

## Development

Requires Go `1.26+` and Windows.

### Run directly

```powershell
go run .
```

**Note:** Use `go run .` (current directory) not `go run main.go` — the latter only compiles `main.go` and misses the other package files.

### Build

```powershell
go build -v -o InputEventHandler.exe
```

## Security notes

- The server listens on `127.0.0.1` only and has no authentication.
- Do not expose this port to untrusted users/processes on the same machine.

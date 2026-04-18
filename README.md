# InputEventHandler

A Windows-only TCP server that accepts JSON commands and simulates keyboard and mouse input using Windows Virtual Key codes. Simple, easy to use, and human-readable protocol.

## Overview

- **JSON Protocol** — Human-readable input format
- **Windows Virtual Keys** — Uses standard VK_* naming convention
- **Tap Only** — Only supports tap actions (press + hold + release)
- **No Binary Encoding** — Easy to debug and test
- **TCP localhost only** — Listens on `127.0.0.1:6767` by default
- **Minimal Dependencies** — No external input libraries; uses Win32 calls via Go's
  builtin syscall (user32.dll) for keyboard/mouse input

## Building

Build with the standard Go toolchain. The produced binary name is up to you;
the examples below use `input_handler.exe` but `go build` (no -o) will produce a
platform-appropriate binary named after the module.

```bash
go build
```

## Running

```bash
.\InputEventHandler.exe
```

Example output (logs may include the client remote address and active connection count):

```
>>> Input Engine Active <<<
Listening on 127.0.0.1:6767 (PID: 0)
```

Note: the PID value is a placeholder in the current implementation and may be 0.

The server runs until it receives an interrupt/termination signal. It accepts
connections on the configured port (default 6767). Each connection is handled
concurrently in its own goroutine and the server performs a graceful shutdown
when requested.

---

## Protocol

The server accepts **line-delimited JSON** input. Each line is a complete JSON object.

### Request Format

```json
{
  "virtual_key": "VK_A",
  "duration": 50
}
```

- **`virtual_key`** (string, required): Windows Virtual Key name (e.g., `VK_A`, `VK_LBUTTON`, `VK_RETURN`)
- **`duration`** (uint16, required): How long to hold the key in milliseconds

### Examples

**Press the 'A' key for 50ms:**
```json
{"virtual_key": "VK_A", "duration": 50}
```

**Left mouse click for 10ms:**
```json
{"virtual_key": "VK_LBUTTON", "duration": 10}
```

**Press Enter for 100ms:**
```json
{"virtual_key": "VK_RETURN", "duration": 100}
```

**Instant press (duration = 0):**
```json
{"virtual_key": "VK_SPACE", "duration": 0}
```

---

## Virtual Key Reference

### Mouse Buttons

| Virtual Key   | Description         |
|---------------|---------------------|
| `VK_LBUTTON`  | Left mouse button   |
| `VK_RBUTTON`  | Right mouse button  |
| `VK_MBUTTON`  | Middle mouse button |
| `VK_XBUTTON1` | X1 mouse button     |
| `VK_XBUTTON2` | X2 mouse button     |

### Keyboard Letters (A-Z)

| Virtual Key           | Description |
|-----------------------|-------------|
| `VK_A` through `VK_Z` | Letter keys |

### Numbers (0-9)

| Virtual Key           | Description |
|-----------------------|-------------|
| `VK_0` through `VK_9` | Number keys |

### Function Keys

| Virtual Key              | Description          |
|--------------------------|----------------------|
| `VK_F1` through `VK_F12` | Function keys F1-F12 |

### Special Keys

| Virtual Key | Description   |
|-------------|---------------|
| `VK_SPACE`  | Spacebar      |
| `VK_RETURN` | Enter key     |
| `VK_ESCAPE` | Escape key    |
| `VK_BACK`   | Backspace key |
| `VK_TAB`    | Tab key       |
| `VK_DELETE` | Delete key    |
| `VK_INSERT` | Insert key    |

### Navigation Keys

| Virtual Key | Description   |
|-------------|---------------|
| `VK_HOME`   | Home key      |
| `VK_END`    | End key       |
| `VK_PRIOR`  | Page Up key   |
| `VK_NEXT`   | Page Down key |

### Arrow Keys

| Virtual Key | Description |
|-------------|-------------|
| `VK_UP`     | Up arrow    |
| `VK_DOWN`   | Down arrow  |
| `VK_LEFT`   | Left arrow  |
| `VK_RIGHT`  | Right arrow |

### Modifier Keys

| Virtual Key   | Description   |
|---------------|---------------|
| `VK_SHIFT`    | Shift key     |
| `VK_LSHIFT`   | Left Shift    |
| `VK_RSHIFT`   | Right Shift   |
| `VK_CONTROL`  | Control key   |
| `VK_LCONTROL` | Left Control  |
| `VK_RCONTROL` | Right Control |
| `VK_MENU`     | Alt key       |
| `VK_LMENU`    | Left Alt      |
| `VK_RMENU`    | Right Alt     |

### Lock Keys

| Virtual Key  | Description |
|--------------|-------------|
| `VK_CAPITAL` | Caps Lock   |
| `VK_NUMLOCK` | Num Lock    |
| `VK_SCROLL`  | Scroll Lock |

### Numpad Keys

| Virtual Key                       | Description |
|-----------------------------------|-------------|
| `VK_NUMPAD0` through `VK_NUMPAD9` | Numpad 0-9  |
| `VK_MULTIPLY`                     | Numpad *    |
| `VK_ADD`                          | Numpad +    |
| `VK_SUBTRACT`                     | Numpad -    |
| `VK_DECIMAL`                      | Numpad .    |
| `VK_DIVIDE`                       | Numpad /    |

For a complete reference, see `keymap/keymap.go` in the source code.

---

## Client Examples

### Go Client (example in repo)

```go
package examples

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func sendInput(conn net.Conn, virtualKey string, duration uint16) error {
	request := map[string]interface{}{
		"virtual_key": virtualKey,
		"duration":    duration,
	}
	jsonData, _ := json.Marshal(request)
	_, err := conn.Write(append(jsonData, '\n'))
	return err
}

func example() {
	conn, err := net.Dial("tcp", "127.0.0.1:6767")
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer conn.Close()

	// Press 'A' for 50ms
	sendInput(conn, "VK_A", 50)
	time.Sleep(100 * time.Millisecond)

	// Left click for 10ms
	sendInput(conn, "VK_LBUTTON", 10)
	time.Sleep(100 * time.Millisecond)

	// Press Enter
	sendInput(conn, "VK_RETURN", 0)
}
```

### Python Client

```python
import socket
import json
import time

def send_input(sock, virtual_key, duration):
    request = {
        "virtual_key": virtual_key,
        "duration": duration
    }
    sock.sendall((json.dumps(request) + '\n').encode())
    print(f"Sent: {request}")

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(("127.0.0.1", 6767))

# Press 'A' for 50ms
send_input(sock, "VK_A", 50)
time.sleep(0.1)

# Left click for 10ms
send_input(sock, "VK_LBUTTON", 10)
time.sleep(0.1)

# Press Enter
send_input(sock, "VK_RETURN", 0)

sock.close()
```

### Java Client

```java
import java.io.*;
import java.net.Socket;
import org.json.JSONObject;

public class InputClient {
    private Socket socket;
    private PrintWriter out;

    public InputClient(String host, int port) throws IOException {
        socket = new Socket(host, port);
        out = new PrintWriter(socket.getOutputStream(), true);
    }

    public void sendInput(String virtualKey, int duration) throws IOException {
        JSONObject request = new JSONObject();
        request.put("virtual_key", virtualKey);
        request.put("duration", duration);
        out.println(request.toString());
        System.out.println("Sent: " + request.toString());
    }

    public void close() throws IOException {
        socket.close();
    }

    public static void example(String[] args) throws IOException, InterruptedException {
        InputClient client = new InputClient("127.0.0.1", 6767);

        // Press 'A' for 50ms
        client.sendInput("VK_A", 50);
        Thread.sleep(100);

        // Left click for 10ms
        client.sendInput("VK_LBUTTON", 10);
        Thread.sleep(100);

        // Press Enter
        client.sendInput("VK_RETURN", 0);

        client.close();
    }
}
```

### JavaScript/Node.js Client

```javascript
const net = require('net');

const socket = net.createConnection({ host: '127.0.0.1', port: 6767 });

function sendInput(virtualKey, duration) {
    const request = {
        virtual_key: virtualKey,
        duration: duration
    };
    socket.write(JSON.stringify(request) + '\n');
    console.log('Sent:', request);
}

socket.on('connect', () => {
    // Press 'A' for 50ms
    sendInput('VK_A', 50);

    setTimeout(() => {
        // Left click for 10ms
        sendInput('VK_LBUTTON', 10);
    }, 100);

    setTimeout(() => {
        // Press Enter
        sendInput('VK_RETURN', 0);
    }, 200);

    setTimeout(() => {
        socket.end();
    }, 300);
});
```

### cURL / Bash

You can use netcat or bash to send a single JSON line to the server:

```bash
# Press 'A' key
echo '{"virtual_key":"VK_A","duration":50}' | nc localhost 6767

# Left click
echo '{"virtual_key":"VK_LBUTTON","duration":10}' | nc localhost 6767
```

---

## Architecture

The project is organized into small Go packages:

- **`main.go`** — program entry point; parses flags and starts the listener
- **`listener/listener.go`** — TCP server, connection handling, JSON parsing
- **`input/input.go`** — Win32 input wrapper (uses user32.dll via syscall)
- **`keymap/keymap.go`** — Virtual Key name → code mapping
- **`examples/client_example.go`** — small Go client showing how to send requests

---

## Error Handling

The server logs activity and errors to stdout. Example lines you may see:

```
Connected: 127.0.0.1:54321 (active: 1)
[127.0.0.1:54321] ✓ VirtualKey=VK_A Duration=50ms
[127.0.0.1:54321] ✓ VirtualKey=VK_LBUTTON Duration=10ms
Disconnected: 127.0.0.1:54321 (clean)
```

If a virtual key is unknown the server logs the error but keeps the connection open:

```
[127.0.0.1:54321] Unknown virtual key: VK_INVALID
```

The server continues processing subsequent lines from the same client after
logging errors; it does not disconnect the client on malformed requests.

---

## Limitations

- **Windows only** — uses Win32 keyboard/mouse simulation
- **Localhost only** — listens on `127.0.0.1` for security
- **Tap only** — only supports press-and-release actions, not separate press/release
- **No authentication** — anyone on your machine with the port can send input

---

## License

MIT

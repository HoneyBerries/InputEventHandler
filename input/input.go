package input

import (
	"InputEventHandler/keymap"
	"syscall"
	"time"
)

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	procMouseEvent    = user32.NewProc("mouse_event")
	procKeyboardEvent = user32.NewProc("keybd_event")
)

const (
	// Mouse event flags
	MouseeventLeftDown = 0x0002
	MouseeventLeftUp   = 0x0004

	MouseeventRightDown = 0x0008
	MouseeventRightUp   = 0x0010

	MouseeventMiddleDown = 0x0020
	MouseeventMiddleUp   = 0x0040

	MouseeventXDown = 0x0080
	MouseeventXUp   = 0x0100

	MouseeventXButton1 = 0x0001
	MouseeventXButton2 = 0x0002
)

// Tap performs a quick press and release with optional hold duration
func Tap(k keymap.KeyDef, durationMs uint16) {
	if k.Code >= 256 {
		// Mouse button (256=left, 257=right, 258=middle, 259=x1, 260=x2)
		tapMouse(k.Code, durationMs)
	} else {
		// Keyboard key (scan code)
		tapKeyboard(k.Code, durationMs)
	}
}

func tapKeyboard(vk uint16, durationMs uint16) {
	// Key down
	procKeyboardEvent.Call(uintptr(vk), 0, 0, 0)

	if durationMs > 0 {
		time.Sleep(time.Duration(durationMs) * time.Millisecond)
	}

	// Key up (0x80 = KEYEVENTF_KEYUP)
	procKeyboardEvent.Call(uintptr(vk), 0, 0x80, 0)
}

func tapMouse(mouseBtn uint16, durationMs uint16) {
	// Press
	switch mouseBtn {
	case 256:
		procMouseEvent.Call(uintptr(MouseeventLeftDown), 0, 0, 0, 0)
	case 257:
		procMouseEvent.Call(uintptr(MouseeventRightDown), 0, 0, 0, 0)
	case 258:
		procMouseEvent.Call(uintptr(MouseeventMiddleDown), 0, 0, 0, 0)
	case 259:
		procMouseEvent.Call(uintptr(MouseeventXDown), 0, 0, uintptr(MouseeventXButton1), 0)
	case 260:
		procMouseEvent.Call(uintptr(MouseeventXDown), 0, 0, uintptr(MouseeventXButton2), 0)
	}

	time.Sleep(time.Duration(durationMs) * time.Millisecond)

	// Release
	switch mouseBtn {
	case 256:
		procMouseEvent.Call(uintptr(MouseeventLeftUp), 0, 0, 0, 0)
	case 257:
		procMouseEvent.Call(uintptr(MouseeventRightUp), 0, 0, 0, 0)
	case 258:
		procMouseEvent.Call(uintptr(MouseeventMiddleUp), 0, 0, 0, 0)
	case 259:
		procMouseEvent.Call(uintptr(MouseeventXUp), 0, 0, uintptr(MouseeventXButton1), 0)
	case 260:
		procMouseEvent.Call(uintptr(MouseeventXUp), 0, 0, uintptr(MouseeventXButton2), 0)
	}
}

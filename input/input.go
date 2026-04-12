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
	MOUSEEVENTF_LEFTDOWN   = 0x0002
	MOUSEEVENTF_LEFTUP     = 0x0004
	MOUSEEVENTF_RIGHTDOWN  = 0x0008
	MOUSEEVENTF_RIGHTUP    = 0x0010
	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP   = 0x0040
)

// Tap performs a quick press and release with optional hold duration
func Tap(k keymap.KeyDef, durationMs uint16) {
	if k.Code >= 256 {
		// Mouse button (256=left, 257=right, 258=middle)
		tapMouse(int(k.Code-256), durationMs)
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

func tapMouse(mouseBtn int, durationMs uint16) {
	// Press
	switch mouseBtn {
	case 1:
		procMouseEvent.Call(uintptr(MOUSEEVENTF_LEFTDOWN), 0, 0, 0, 0)
	case 2:
		procMouseEvent.Call(uintptr(MOUSEEVENTF_RIGHTDOWN), 0, 0, 0, 0)
	case 3:
		procMouseEvent.Call(uintptr(MOUSEEVENTF_MIDDLEDOWN), 0, 0, 0, 0)
	}

	if durationMs > 0 {
		time.Sleep(time.Duration(durationMs) * time.Millisecond)
	}

	// Release
	switch mouseBtn {
	case 1:
		procMouseEvent.Call(uintptr(MOUSEEVENTF_LEFTUP), 0, 0, 0, 0)
	case 2:
		procMouseEvent.Call(uintptr(MOUSEEVENTF_RIGHTUP), 0, 0, 0, 0)
	case 3:
		procMouseEvent.Call(uintptr(MOUSEEVENTF_MIDDLEUP), 0, 0, 0, 0)
	}
}

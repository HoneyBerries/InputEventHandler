package main

import (
	"fmt"
	"syscall"
	"time"
)

var (
	user32    = syscall.NewLazyDLL("user32.dll")
	mouseProc = user32.NewProc("mouse_event")
	kbdProc   = user32.NewProc("keybd_event")
)

const keyUp = uintptr(0x80)

const (
	mouseLeft   uint16 = 256
	mouseRight  uint16 = 257
	mouseMiddle uint16 = 258
	mouseX1     uint16 = 259
	mouseX2     uint16 = 260
)

type mouseBtn struct{ down, up, xdata uint16 }

var mouseBtns = map[uint16]mouseBtn{
	mouseLeft:   {0x0002, 0x0004, 0},
	mouseRight:  {0x0008, 0x0010, 0},
	mouseMiddle: {0x0020, 0x0040, 0},
	mouseX1:     {0x0080, 0x0100, 0x0001},
	mouseX2:     {0x0080, 0x0100, 0x0002},
}

func tap(keyCode, durationMs uint16) error {
	if keyCode >= mouseLeft {
		return tapMouse(keyCode, durationMs)
	}
	return tapKeyboard(keyCode, durationMs)
}

func tapKeyboard(vk, durationMs uint16) error {
	if _, _, err := kbdProc.Call(uintptr(vk), 0, 0, 0); !winOK(err) {
		return fmt.Errorf("keydown %#x: %w", vk, err)
	}
	time.Sleep(time.Duration(durationMs) * time.Millisecond)
	if _, _, err := kbdProc.Call(uintptr(vk), 0, keyUp, 0); !winOK(err) {
		return fmt.Errorf("keyup %#x: %w", vk, err)
	}
	return nil
}

func tapMouse(btn, durationMs uint16) error {
	b, ok := mouseBtns[btn]
	if !ok {
		return fmt.Errorf("unknown mouse button: %d", btn)
	}
	if _, _, err := mouseProc.Call(uintptr(b.down), 0, 0, uintptr(b.xdata), 0); !winOK(err) {
		return fmt.Errorf("mousedown %d: %w", btn, err)
	}
	time.Sleep(time.Duration(durationMs) * time.Millisecond)
	if _, _, err := mouseProc.Call(uintptr(b.up), 0, 0, uintptr(b.xdata), 0); !winOK(err) {
		return fmt.Errorf("mouseup %d: %w", btn, err)
	}
	return nil
}

// winOK returns true when the syscall "error" is Windows' success sentinel.
// syscall.LazyProc.Call always populates err from GetLastError(); on success
// that is ERROR_SUCCESS (0), which Go maps to this fixed string.
func winOK(err error) bool {
	return err.Error() == "The operation completed successfully."
}

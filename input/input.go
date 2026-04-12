package input

import (
	"fmt"
	"syscall"
	"time"
)

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	procMouseEvent    = user32.NewProc("mouse_event")
	procKeyboardEvent = user32.NewProc("keybd_event")
)

// Keyboard scan codes
const (
	KeyEventKeyUp = 0x80
)

// Mouse button codes (virtual key codes 256+)
const (
	MouseButtonLeft   = 256
	MouseButtonRight  = 257
	MouseButtonMiddle = 258
	MouseButtonX1     = 259
	MouseButtonX2     = 260
)

// Mouse event flags
const (
	MouseEventLeftDown   = 0x0002
	MouseEventLeftUp     = 0x0004
	MouseEventRightDown  = 0x0008
	MouseEventRightUp    = 0x0010
	MouseEventMiddleDown = 0x0020
	MouseEventMiddleUp   = 0x0040
	MouseEventXDown      = 0x0080
	MouseEventXUp        = 0x0100
	MouseEventXButton1   = 0x0001
	MouseEventXButton2   = 0x0002
)

// MouseButtonInfo maps button codes to their event flags and x-button identifier
type MouseButtonInfo struct {
	DownFlag uint16
	UpFlag   uint16
	XButton  uint16 // 0 if not an X button, else MouseEventXButton1/2
}

// mouseButtonMap provides O(1) lookup for mouse button event flags
var mouseButtonMap = map[uint16]MouseButtonInfo{
	MouseButtonLeft: {
		DownFlag: MouseEventLeftDown,
		UpFlag:   MouseEventLeftUp,
		XButton:  0,
	},
	MouseButtonRight: {
		DownFlag: MouseEventRightDown,
		UpFlag:   MouseEventRightUp,
		XButton:  0,
	},
	MouseButtonMiddle: {
		DownFlag: MouseEventMiddleDown,
		UpFlag:   MouseEventMiddleUp,
		XButton:  0,
	},
	MouseButtonX1: {
		DownFlag: MouseEventXDown,
		UpFlag:   MouseEventXUp,
		XButton:  MouseEventXButton1,
	},
	MouseButtonX2: {
		DownFlag: MouseEventXDown,
		UpFlag:   MouseEventXUp,
		XButton:  MouseEventXButton2,
	},
}

// Tap performs a quick press and release of a key with optional hold duration.
// keyCode should be a virtual key code (0-255 for keyboard, 256+ for mouse).
// durationMs specifies how long to hold the key (0 = press and release immediately).
// Returns an error if the syscall fails.
func Tap(keyCode uint16, durationMs uint16) error {
	if keyCode >= MouseButtonLeft {
		return tapMouse(keyCode, durationMs)
	}
	return tapKeyboard(keyCode, durationMs)
}

// tapKeyboard sends a keyboard key press/release sequence.
func tapKeyboard(vk uint16, durationMs uint16) error {
	// Key press
	ret, _, err := procKeyboardEvent.Call(uintptr(vk), 0, 0, 0)
	if ret == 0 && err != nil && !isSuccessError(err) {
		return fmt.Errorf("keyboard down event failed: %w", err)
	}

	time.Sleep(time.Duration(durationMs) * time.Millisecond)

	// Key release
	ret, _, err = procKeyboardEvent.Call(uintptr(vk), 0, KeyEventKeyUp, 0)
	if ret == 0 && err != nil && !isSuccessError(err) {
		return fmt.Errorf("keyboard up event failed: %w", err)
	}

	return nil
}

// tapMouse sends a mouse button press/release sequence.
func tapMouse(mouseBtn uint16, durationMs uint16) error {
	btnInfo, exists := mouseButtonMap[mouseBtn]
	if !exists {
		return fmt.Errorf("invalid mouse button code: %d (valid: 256-260)", mouseBtn)
	}

	// Press
	if err := sendMouseEvent(btnInfo.DownFlag, btnInfo.XButton); err != nil {
		return fmt.Errorf("mouse down event (button %d) failed: %w", mouseBtn, err)
	}

	time.Sleep(time.Duration(durationMs) * time.Millisecond)

	// Release
	if err := sendMouseEvent(btnInfo.UpFlag, btnInfo.XButton); err != nil {
		return fmt.Errorf("mouse up event (button %d) failed: %w", mouseBtn, err)
	}

	return nil
}

// sendMouseEvent is a helper that dispatches a single mouse event to the system.
func sendMouseEvent(eventFlag uint16, xButton uint16) error {
	ret, _, err := procMouseEvent.Call(uintptr(eventFlag), 0, 0, uintptr(xButton), 0)
	if ret == 0 && err != nil && !isSuccessError(err) {
		return err
	}
	return nil
}

// isSuccessError checks if a syscall error is actually a success (Windows returns 0 = success as error)
func isSuccessError(err error) bool {
	return err.Error() == "The operation completed successfully."
}

package keymap

// KeyCode is a type alias for key codes to improve type safety
type KeyCode uint16

// KeyDef defines a keyboard/mouse key
// Keyboard: Code = virtual key (0x01 - 0xFF)
// Mouse: Code = 256 (left), 257 (right), 258 (middle), 259 (x1), 260 (x2)
type KeyDef struct {
	Code KeyCode
}

// VirtualKeyMap maps Windows virtual key names to their codes.
// Keyboard codes are 0x01-0xFF (standard Windows virtual key codes).
// Mouse codes are 256-260 (custom encoding for safe distinction from keyboard).
var VirtualKeyMap = map[string]KeyCode{
	// Mouse buttons (256+ for safe distinction from keyboard keys)
	"VK_LBUTTON":  256,
	"VK_RBUTTON":  257,
	"VK_MBUTTON":  258,
	"VK_XBUTTON1": 259,
	"VK_XBUTTON2": 260,

	// Function keys
	"VK_F1":  0x70,
	"VK_F2":  0x71,
	"VK_F3":  0x72,
	"VK_F4":  0x73,
	"VK_F5":  0x74,
	"VK_F6":  0x75,
	"VK_F7":  0x76,
	"VK_F8":  0x77,
	"VK_F9":  0x78,
	"VK_F10": 0x79,
	"VK_F11": 0x7A,
	"VK_F12": 0x7B,

	// Alphanumeric keys (A-Z, 0-9)
	"VK_0": 0x30,
	"VK_1": 0x31,
	"VK_2": 0x32,
	"VK_3": 0x33,
	"VK_4": 0x34,
	"VK_5": 0x35,
	"VK_6": 0x36,
	"VK_7": 0x37,
	"VK_8": 0x38,
	"VK_9": 0x39,

	"VK_A": 0x41,
	"VK_B": 0x42,
	"VK_C": 0x43,
	"VK_D": 0x44,
	"VK_E": 0x45,
	"VK_F": 0x46,
	"VK_G": 0x47,
	"VK_H": 0x48,
	"VK_I": 0x49,
	"VK_J": 0x4A,
	"VK_K": 0x4B,
	"VK_L": 0x4C,
	"VK_M": 0x4D,
	"VK_N": 0x4E,
	"VK_O": 0x4F,
	"VK_P": 0x50,
	"VK_Q": 0x51,
	"VK_R": 0x52,
	"VK_S": 0x53,
	"VK_T": 0x54,
	"VK_U": 0x55,
	"VK_V": 0x56,
	"VK_W": 0x57,
	"VK_X": 0x58,
	"VK_Y": 0x59,
	"VK_Z": 0x5A,

	// Special keys
	"VK_BACK":   0x08,
	"VK_TAB":    0x09,
	"VK_RETURN": 0x0D,
	"VK_ESCAPE": 0x1B,
	"VK_SPACE":  0x20,
	"VK_DELETE": 0x2E,
	"VK_INSERT": 0x2D,

	// Navigation keys
	"VK_HOME":  0x24,
	"VK_END":   0x23,
	"VK_PRIOR": 0x21, // Page Up
	"VK_NEXT":  0x22, // Page Down

	// Arrow keys
	"VK_LEFT":  0x25,
	"VK_RIGHT": 0x27,
	"VK_UP":    0x26,
	"VK_DOWN":  0x28,

	// Modifier keys
	"VK_SHIFT":    0x10,
	"VK_CONTROL":  0x11,
	"VK_MENU":     0x12, // Alt
	"VK_LSHIFT":   0xA0,
	"VK_RSHIFT":   0xA1,
	"VK_LCONTROL": 0xA2,
	"VK_RCONTROL": 0xA3,
	"VK_LMENU":    0xA4, // Left Alt
	"VK_RMENU":    0xA5, // Right Alt

	// Lock keys
	"VK_CAPITAL": 0x14,
	"VK_NUMLOCK": 0x90,
	"VK_SCROLL":  0x91,

	// Numpad keys
	"VK_NUMPAD0":  0x60,
	"VK_NUMPAD1":  0x61,
	"VK_NUMPAD2":  0x62,
	"VK_NUMPAD3":  0x63,
	"VK_NUMPAD4":  0x64,
	"VK_NUMPAD5":  0x65,
	"VK_NUMPAD6":  0x66,
	"VK_NUMPAD7":  0x67,
	"VK_NUMPAD8":  0x68,
	"VK_NUMPAD9":  0x69,
	"VK_MULTIPLY": 0x6A,
	"VK_ADD":      0x6B,
	"VK_SUBTRACT": 0x6D,
	"VK_DECIMAL":  0x6E,
	"VK_DIVIDE":   0x6F,

	// Additional keys
	"VK_PAUSE":    0x13,
	"VK_SNAPSHOT": 0x2C, // Print Screen
	"VK_PRINT":    0x2A,
}

// KeyCodeByName looks up a key code by name.
// Returns (keyCode, ok) where ok is true if the key name was found.
func KeyCodeByName(name string) (KeyCode, bool) {
	code, ok := VirtualKeyMap[name]
	return code, ok
}

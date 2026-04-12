package keymap

// KeyDef defines a keyboard/mouse key
// Keyboard: Code = virtual key (0x01 - 0xFF)
// Mouse: Code = 256 (left), 257 (right), 258 (middle)
type KeyDef struct {
	Code uint16 // virtual key for keyboard, or 256+ for mouse buttons
}

// VirtualKeyMap maps Windows virtual key names to KeyDef
// Stores Windows Virtual Key codes, not scan codes
var VirtualKeyMap = map[string]KeyDef{
	// Mouse buttons (256+ for safe distinction)
	"VK_LBUTTON":  {Code: 256},
	"VK_RBUTTON":  {Code: 257},
	"VK_MBUTTON":  {Code: 258},
	"VK_XBUTTON1": {Code: 256},
	"VK_XBUTTON2": {Code: 257},

	// Function keys
	"VK_F1":  {Code: 0x70},
	"VK_F2":  {Code: 0x71},
	"VK_F3":  {Code: 0x72},
	"VK_F4":  {Code: 0x73},
	"VK_F5":  {Code: 0x74},
	"VK_F6":  {Code: 0x75},
	"VK_F7":  {Code: 0x76},
	"VK_F8":  {Code: 0x77},
	"VK_F9":  {Code: 0x78},
	"VK_F10": {Code: 0x79},
	"VK_F11": {Code: 0x7A},
	"VK_F12": {Code: 0x7B},

	// Alphanumeric keys (A-Z, 0-9)
	"VK_0": {Code: 0x30},
	"VK_1": {Code: 0x31},
	"VK_2": {Code: 0x32},
	"VK_3": {Code: 0x33},
	"VK_4": {Code: 0x34},
	"VK_5": {Code: 0x35},
	"VK_6": {Code: 0x36},
	"VK_7": {Code: 0x37},
	"VK_8": {Code: 0x38},
	"VK_9": {Code: 0x39},

	"VK_A": {Code: 0x41},
	"VK_B": {Code: 0x42},
	"VK_C": {Code: 0x43},
	"VK_D": {Code: 0x44},
	"VK_E": {Code: 0x45},
	"VK_F": {Code: 0x46},
	"VK_G": {Code: 0x47},
	"VK_H": {Code: 0x48},
	"VK_I": {Code: 0x49},
	"VK_J": {Code: 0x4A},
	"VK_K": {Code: 0x4B},
	"VK_L": {Code: 0x4C},
	"VK_M": {Code: 0x4D},
	"VK_N": {Code: 0x4E},
	"VK_O": {Code: 0x4F},
	"VK_P": {Code: 0x50},
	"VK_Q": {Code: 0x51},
	"VK_R": {Code: 0x52},
	"VK_S": {Code: 0x53},
	"VK_T": {Code: 0x54},
	"VK_U": {Code: 0x55},
	"VK_V": {Code: 0x56},
	"VK_W": {Code: 0x57},
	"VK_X": {Code: 0x58},
	"VK_Y": {Code: 0x59},
	"VK_Z": {Code: 0x5A},

	// Special keys
	"VK_BACK":   {Code: 0x08},
	"VK_TAB":    {Code: 0x09},
	"VK_RETURN": {Code: 0x0D},
	"VK_ESCAPE": {Code: 0x1B},
	"VK_SPACE":  {Code: 0x20},
	"VK_DELETE": {Code: 0x2E},
	"VK_INSERT": {Code: 0x2D},

	// Navigation keys
	"VK_HOME":  {Code: 0x24},
	"VK_END":   {Code: 0x23},
	"VK_PRIOR": {Code: 0x21}, // Page Up
	"VK_NEXT":  {Code: 0x22}, // Page Down

	// Arrow keys
	"VK_LEFT":  {Code: 0x25},
	"VK_RIGHT": {Code: 0x27},
	"VK_UP":    {Code: 0x26},
	"VK_DOWN":  {Code: 0x28},

	// Modifier keys
	"VK_SHIFT":    {Code: 0x10},
	"VK_CONTROL":  {Code: 0x11},
	"VK_MENU":     {Code: 0x12}, // Alt
	"VK_LSHIFT":   {Code: 0xA0},
	"VK_RSHIFT":   {Code: 0xA1},
	"VK_LCONTROL": {Code: 0xA2},
	"VK_RCONTROL": {Code: 0xA3},
	"VK_LMENU":    {Code: 0xA4},
	"VK_RMENU":    {Code: 0xA5},

	// Lock keys
	"VK_CAPITAL": {Code: 0x14},
	"VK_NUMLOCK": {Code: 0x90},
	"VK_SCROLL":  {Code: 0x91},

	// Numpad keys
	"VK_NUMPAD0":  {Code: 0x60},
	"VK_NUMPAD1":  {Code: 0x61},
	"VK_NUMPAD2":  {Code: 0x62},
	"VK_NUMPAD3":  {Code: 0x63},
	"VK_NUMPAD4":  {Code: 0x64},
	"VK_NUMPAD5":  {Code: 0x65},
	"VK_NUMPAD6":  {Code: 0x66},
	"VK_NUMPAD7":  {Code: 0x67},
	"VK_NUMPAD8":  {Code: 0x68},
	"VK_NUMPAD9":  {Code: 0x69},
	"VK_MULTIPLY": {Code: 0x6A},
	"VK_ADD":      {Code: 0x6B},
	"VK_SUBTRACT": {Code: 0x6D},
	"VK_DECIMAL":  {Code: 0x6E},
	"VK_DIVIDE":   {Code: 0x6F},

	// Additional keys
	"VK_PAUSE":    {Code: 0x13},
	"VK_SNAPSHOT": {Code: 0x2C}, // Print Screen
	"VK_PRINT":    {Code: 0x2A},
}

package internal

/*
#cgo LDFLAGS: -framework ApplicationServices -framework CoreFoundation
#include "hotkey_darwin.h"
*/
import "C"

import (
	"fmt"
)

const (
	KeyRightOption  = 61
	KeyRightCommand = 54
	KeyFn           = 63
	KeyF18          = 79
	KeyF19          = 80
)

type HotkeyInfo struct {
	KeyCode C.CGKeyCode
	ModFlag C.CGEventFlags
}

var (
	onDown func()
	onUp   func()
)

//export goHotkeyDown
func goHotkeyDown() {
	if onDown != nil {
		onDown()
	}
}

//export goHotkeyUp
func goHotkeyUp() {
	if onUp != nil {
		onUp()
	}
}

func ResolveHotkey(name string) (HotkeyInfo, error) {
	switch name {
	case "right_option", "right_alt":
		return HotkeyInfo{
			KeyCode: C.CGKeyCode(KeyRightOption),
			ModFlag: C.CGEventFlags(C.kCGEventFlagMaskAlternate),
		}, nil
	case "right_command", "right_cmd":
		return HotkeyInfo{
			KeyCode: C.CGKeyCode(KeyRightCommand),
			ModFlag: C.CGEventFlags(C.kCGEventFlagMaskCommand),
		}, nil
	case "fn":
		return HotkeyInfo{
			KeyCode: C.CGKeyCode(KeyFn),
			ModFlag: C.CGEventFlags(C.kCGEventFlagMaskSecondaryFn),
		}, nil
	case "f18":
		return HotkeyInfo{KeyCode: C.CGKeyCode(KeyF18)}, nil
	case "f19":
		return HotkeyInfo{KeyCode: C.CGKeyCode(KeyF19)}, nil
	default:
		return HotkeyInfo{}, fmt.Errorf("unknown hotkey: %s (supported: right_option, right_command, fn, f18, f19)", name)
	}
}

func CheckAccessibility() bool {
	return C.AXIsProcessTrusted() != 0
}

func ListenHotkey(hk HotkeyInfo, downFn, upFn func()) error {
	onDown = downFn
	onUp = upFn

	if !C.startEventTap(hk.KeyCode, hk.ModFlag) {
		return fmt.Errorf("failed to create event tap — check Accessibility permissions")
	}

	C.runLoop()
	return nil
}

func StopHotkey() {
	C.stopLoop()
}

package internal

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa -framework QuartzCore
// #include "overlay_darwin.h"
import "C"

var overlayEnabled bool

func OverlayInit(enabled bool) {
	overlayEnabled = enabled
	if !overlayEnabled {
		return
	}
	C.overlayInit()
}

func OverlayShow(state int) {
	if !overlayEnabled {
		return
	}
	C.overlayShow(C.int(state))
}

func OverlayHide() {
	if !overlayEnabled {
		return
	}
	C.overlayHide()
}

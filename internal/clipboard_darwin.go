package internal

/*
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>
#include <unistd.h>

void simulatePaste() {
    CGEventRef down = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)9, true);  // Cmd+V
    CGEventSetFlags(down, kCGEventFlagMaskCommand);
    CGEventPost(kCGHIDEventTap, down);
    CFRelease(down);
    usleep(5000);

    CGEventRef up = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)9, false);
    CGEventSetFlags(up, kCGEventFlagMaskCommand);
    CGEventPost(kCGHIDEventTap, up);
    CFRelease(up);
}
*/
import "C"

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
)

type ClipboardMode struct{}

func (c *ClipboardMode) ShowLoading() {
	// No-op — terminal inputs don't support reliable text replacement
}

func (c *ClipboardMode) Deliver(text string) error {
	fmt.Printf("[DEBUG clipboard] Deliver called with text=%q\n", text)

	if err := clipboard.WriteAll(text); err != nil {
		fmt.Printf("[DEBUG clipboard] WriteAll error: %v\n", err)
		return err
	}

	check, _ := clipboard.ReadAll()
	fmt.Printf("[DEBUG clipboard] clipboard now=%q\n", check)

	time.Sleep(50 * time.Millisecond)

	fmt.Println("[DEBUG clipboard] simulating Cmd+V paste...")
	C.simulatePaste()

	return nil
}

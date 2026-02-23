#include "hotkey_darwin.h"
#include "_cgo_export.h"

static CGKeyCode targetKeyCode = 0;
static bool keyIsDown = false;
static CGEventFlags targetFlag = 0;
static CFMachPortRef tapRef = NULL;
static CFRunLoopRef mainRunLoop = NULL;

static CGEventRef eventCallback(CGEventTapProxy proxy, CGEventType evType, CGEventRef event, void *refcon) {
    (void)proxy;
    (void)refcon;

    // Handle tap being disabled by the system
    if (evType == (CGEventType)0xFFFFFFFE) {
        if (tapRef != NULL) {
            CGEventTapEnable(tapRef, true);
        }
        return event;
    }

    // For modifier keys (Option, Command, Fn), use FlagsChanged events
    if (targetFlag != 0) {
        if (evType == kCGEventFlagsChanged) {
            CGEventFlags flags = CGEventGetFlags(event);
            bool isPressed = (flags & targetFlag) != 0;

            if (isPressed && !keyIsDown) {
                keyIsDown = true;
                goHotkeyDown();
                return NULL;
            } else if (!isPressed && keyIsDown) {
                keyIsDown = false;
                goHotkeyUp();
                return NULL;
            }
        }
        return event;
    }

    // For regular keys, use keyDown/keyUp
    CGKeyCode keyCode = (CGKeyCode)CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);
    if (keyCode != targetKeyCode) {
        return event;
    }

    if (evType == kCGEventKeyDown) {
        if (!keyIsDown) {
            keyIsDown = true;
            goHotkeyDown();
        }
        return NULL;
    }

    if (evType == kCGEventKeyUp) {
        if (keyIsDown) {
            keyIsDown = false;
            goHotkeyUp();
        }
        return NULL;
    }

    return event;
}

bool startEventTap(CGKeyCode keyCode, CGEventFlags modFlag) {
    targetKeyCode = keyCode;
    targetFlag = modFlag;
    keyIsDown = false;

    CGEventMask mask = CGEventMaskBit(kCGEventKeyDown) |
                       CGEventMaskBit(kCGEventKeyUp) |
                       CGEventMaskBit(kCGEventFlagsChanged);

    tapRef = CGEventTapCreate(
        kCGSessionEventTap,
        kCGHeadInsertEventTap,
        kCGEventTapOptionDefault,
        mask,
        eventCallback,
        NULL
    );

    if (tapRef == NULL) {
        return false;
    }

    CFRunLoopSourceRef source = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, tapRef, 0);
    CFRunLoopAddSource(CFRunLoopGetCurrent(), source, kCFRunLoopCommonModes);
    CGEventTapEnable(tapRef, true);
    CFRelease(source);

    return true;
}

void runLoop(void) {
    mainRunLoop = CFRunLoopGetCurrent();
    CFRunLoopRun();
}

void stopLoop(void) {
    if (tapRef != NULL) {
        CGEventTapEnable(tapRef, false);
        CFRelease(tapRef);
        tapRef = NULL;
    }
    if (mainRunLoop != NULL) {
        CFRunLoopStop(mainRunLoop);
    }
}

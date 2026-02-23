#ifndef HOTKEY_DARWIN_H
#define HOTKEY_DARWIN_H

#include <ApplicationServices/ApplicationServices.h>
#include <CoreFoundation/CoreFoundation.h>
#include <stdbool.h>

bool startEventTap(CGKeyCode keyCode, CGEventFlags modFlag);
void runLoop(void);
void stopLoop(void);

#endif

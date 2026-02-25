#import <Cocoa/Cocoa.h>
#import <QuartzCore/QuartzCore.h>

// Embedded mascot PNG data
#include "mascot_data.h"

static const CGFloat kCircleSize = 48.0;
static const CGFloat kBorderWidth = 2.5;
static const CGFloat kBottomMargin = 48.0;
// Total window size accounts for the glow shadow extending beyond the circle
static const CGFloat kShadowPad = 10.0;
static const CGFloat kWindowSize = 48.0 + 20.0; // kCircleSize + 2*kShadowPad

static NSPanel *overlayPanel = nil;
static NSView *ringView = nil;
static NSImageView *imageView = nil;
static bool overlayReady = false;

static void createPanel(void) {
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyAccessory];

    NSRect winFrame = NSMakeRect(0, 0, kWindowSize, kWindowSize);

    overlayPanel = [[NSPanel alloc]
        initWithContentRect:winFrame
                  styleMask:NSWindowStyleMaskBorderless | NSWindowStyleMaskNonactivatingPanel
                    backing:NSBackingStoreBuffered
                      defer:NO];

    [overlayPanel setLevel:NSStatusWindowLevel];
    [overlayPanel setOpaque:NO];
    [overlayPanel setBackgroundColor:[NSColor clearColor]];
    [overlayPanel setIgnoresMouseEvents:YES];
    [overlayPanel setHasShadow:NO];
    [overlayPanel setCollectionBehavior:
        NSWindowCollectionBehaviorCanJoinAllSpaces |
        NSWindowCollectionBehaviorStationary |
        NSWindowCollectionBehaviorFullScreenAuxiliary];

    // Ring view (colored border circle)
    NSRect ringFrame = NSMakeRect(kShadowPad, kShadowPad, kCircleSize, kCircleSize);
    ringView = [[NSView alloc] initWithFrame:ringFrame];
    [ringView setWantsLayer:YES];
    ringView.layer.cornerRadius = kCircleSize / 2.0;
    ringView.layer.borderWidth = kBorderWidth;
    ringView.layer.borderColor = [[NSColor colorWithSRGBRed:0.2 green:0.84 blue:0.4 alpha:1.0] CGColor];
    ringView.layer.backgroundColor = [[NSColor whiteColor] CGColor];
    ringView.layer.masksToBounds = YES;

    // Load embedded mascot image
    NSData *imgData = [NSData dataWithBytesNoCopy:mascot_png length:mascot_png_len freeWhenDone:NO];
    NSImage *mascot = [[NSImage alloc] initWithData:imgData];

    // Image view inside the ring, inset well within the border
    CGFloat inset = kBorderWidth + 4.0;
    CGFloat imgSize = kCircleSize - inset * 2;
    NSRect imgFrame = NSMakeRect(inset, inset, imgSize, imgSize);
    imageView = [[NSImageView alloc] initWithFrame:imgFrame];
    [imageView setImage:mascot];
    [imageView setImageScaling:NSImageScaleProportionallyUpOrDown];
    [imageView setWantsLayer:YES];
    imageView.layer.cornerRadius = imgSize / 2.0;
    imageView.layer.masksToBounds = YES;

    [ringView addSubview:imageView];
    [overlayPanel.contentView addSubview:ringView];

    // Position: bottom center, 10px above screen edge
    NSScreen *screen = [NSScreen mainScreen];
    if (screen) {
        NSRect sv = [screen frame];
        CGFloat x = (sv.size.width - kWindowSize) / 2.0;
        CGFloat y = sv.origin.y + kBottomMargin;
        [overlayPanel setFrameOrigin:NSMakePoint(x, y)];
    }

    overlayReady = true;
}

static void startPulse(void) {
    [ringView.layer removeAnimationForKey:@"glow"];

    // Pulsing glow via shadow
    ringView.layer.shadowColor = ringView.layer.borderColor;
    ringView.layer.shadowOffset = CGSizeMake(0, 0);
    ringView.layer.shadowRadius = 8;
    ringView.layer.shadowOpacity = 0.9;
    ringView.layer.masksToBounds = NO;

    CABasicAnimation *glow = [CABasicAnimation animationWithKeyPath:@"shadowOpacity"];
    glow.fromValue = @0.9;
    glow.toValue = @0.2;
    glow.duration = 0.8;
    glow.autoreverses = YES;
    glow.repeatCount = HUGE_VALF;
    glow.timingFunction = [CAMediaTimingFunction functionWithName:kCAMediaTimingFunctionEaseInEaseOut];
    [ringView.layer addAnimation:glow forKey:@"glow"];
}

static void stopPulse(void) {
    [ringView.layer removeAnimationForKey:@"glow"];
    ringView.layer.shadowOpacity = 0;
}

void overlayInit(void) {
    dispatch_async(dispatch_get_main_queue(), ^{
        createPanel();
    });
}

void overlayShow(int state) {
    dispatch_async(dispatch_get_main_queue(), ^{
        if (!overlayReady) {
            createPanel();
        }

        NSColor *color;
        if (state == 0) {
            // Listening — green ring
            color = [NSColor colorWithSRGBRed:0.2 green:0.84 blue:0.4 alpha:1.0];
        } else {
            // Processing — amber ring
            color = [NSColor colorWithSRGBRed:1.0 green:0.76 blue:0.0 alpha:1.0];
        }
        ringView.layer.borderColor = [color CGColor];

        // Need to temporarily allow overflow for the glow shadow
        ringView.layer.masksToBounds = NO;
        startPulse();

        [overlayPanel orderFrontRegardless];
    });
}

void overlayHide(void) {
    dispatch_async(dispatch_get_main_queue(), ^{
        if (overlayReady) {
            stopPulse();
            [overlayPanel orderOut:nil];
        }
    });
}

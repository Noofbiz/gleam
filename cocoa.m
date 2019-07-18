// +build darwin
// +build 386 amd64
// +build !ios

#import <Cocoa/Cocoa.h>

#include "_cgo_export.h"
#include <float.h>
#include <string.h>

@interface AppDelegate : NSObject<NSApplicationDelegate>
{
}
@end

@implementation AppDelegate
- (void)applicationWillFinishLaunching:(NSNotification *)notification
{
  [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
  gleamInitDoneSignal();
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
}

- (void)applicationWillHide:(NSNotification *)aNotification {
}
@end

@interface WindowDelegate : NSObject<NSWindowDelegate>
{
}
@end

@implementation WindowDelegate
- (BOOL)windowShouldClose:(id)sender
{
    NSLog(@"window close clicked");
    return NO;
}
- (void)windowDidResize:(NSNotification *)notification
{
}

- (void)windowDidMove:(NSNotification *)notification
{
}

- (void)windowDidMiniaturize:(NSNotification *)notification
{
}

- (void)windowDidDeminiaturize:(NSNotification *)notification
{
}

- (void)windowDidBecomeKey:(NSNotification *)notification
{
}

- (void)windowDidResignKey:(NSNotification *)notification
{
}

- (void)windowDidChangeScreen:(NSNotification *)notification
{
}
@end

int initializeGleam() {
  NSAutoreleasePool * pool = [[NSAutoreleasePool alloc] init];
  [NSApplication sharedApplication];

  AppDelegate *appDelegate = [[AppDelegate alloc] init];
  [NSApp setDelegate:appDelegate];
  [NSApp run];
  [pool release];
  return 0;
}

int appDone() {
  dispatch_async(dispatch_get_main_queue(), ^{
		[NSApp terminate:nil];
	});
  return 0;
}

int setTitle(uintptr_t w, char* title) {
  dispatch_sync(dispatch_get_main_queue(), ^{
    NSString* name = [[NSString alloc] initWithUTF8String:title];
    NSWindow* window = (NSWindow*)w;
    window.title = name;
  });
  return 0;
}

int setFullScreen(uintptr_t w) {
  NSScreen *screen = [NSScreen mainScreen];
  NSUInteger mask = NSWindowStyleMaskBorderless;
  NSRect frame = screen.frame;
  dispatch_sync(dispatch_get_main_queue(), ^{
    NSWindow* window = (NSWindow*)w;
    [window setStyleMask: mask];
    [window setFrame: frame display: YES animate: NO];
  });
  return 0;
}

int resize(uintptr_t w, int width, int height, int x, int y, bool resizable) {
  NSScreen *screen = [NSScreen mainScreen];
	double wid = (double)width / [screen backingScaleFactor];
	double h = (double)height / [screen backingScaleFactor];
  __block NSUInteger mask = NSWindowStyleMaskTitled|NSWindowStyleMaskClosable|NSWindowStyleMaskMiniaturizable;
  dispatch_sync(dispatch_get_main_queue(), ^{
    NSWindow* window = (NSWindow*)w;
    NSRect frame = NSMakeRect(0, 0, wid, h);
    [window setStyleMask: mask];
    [window setFrame: frame display: YES animate: NO];
    [window cascadeTopLeftFromPoint:NSMakePoint(x,y)];
    if (resizable) {
      mask |= NSWindowStyleMaskResizable;
      [window setStyleMask: mask];
    }
  });
  return 0;
}

int closeWindow(uintptr_t w) {
  dispatch_sync(dispatch_get_main_queue(), ^{
    NSWindow* window = (NSWindow*)w;
    [window close];
  });
  return 0;
}

uintptr_t newWindow(int width, int height, int x, int y, char* title, bool resizable, bool fullscreen) {
  WindowDelegate *delegate = [[WindowDelegate alloc] init];
  NSScreen *screen = [NSScreen mainScreen];
	double w = (double)width / [screen backingScaleFactor];
	double h = (double)height / [screen backingScaleFactor];
  NSUInteger mask = NSWindowStyleMaskTitled|NSWindowStyleMaskClosable|NSWindowStyleMaskMiniaturizable;
  if (resizable)
    mask |= NSWindowStyleMaskResizable;
  if (fullscreen)
    mask = NSWindowStyleMaskBorderless;
  __block NSWindow* window = NULL;
  dispatch_sync(dispatch_get_main_queue(), ^{
    NSString* name = [[NSString alloc] initWithUTF8String:title];
    NSRect frame = NSMakeRect(0, 0, w, h);
    if (fullscreen)
      frame = screen.frame;
    window  = [[[NSWindow alloc] initWithContentRect:frame
                      styleMask:mask
                      backing:NSBackingStoreBuffered
                      defer:NO] autorelease];
    [window setBackgroundColor:[NSColor blueColor]];
    window.title = name;
    [window setDelegate:delegate];
    if (!fullscreen)
      [window cascadeTopLeftFromPoint:NSMakePoint(x+1,y+1)];
    [NSApp activateIgnoringOtherApps:YES];
    [window makeKeyAndOrderFront:NSApp];
  });
  return (uintptr_t)window;
}

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

uintptr_t newWindow(int width, int height, char* title, bool titled, bool bordered,
bool closable, bool miniaturizable, bool resizable, bool fullscreen) {
  WindowDelegate *delegate = [[WindowDelegate alloc] init];
  NSScreen *screen = [NSScreen mainScreen];
	double w = (double)width / [screen backingScaleFactor];
	double h = (double)height / [screen backingScaleFactor];
  NSUInteger mask = 0;
  if (!bordered)
    mask |= NSWindowStyleMaskBorderless;
  if (titled)
    mask |= NSWindowStyleMaskTitled;
  if (closable)
    mask |= NSWindowStyleMaskClosable;
  if (miniaturizable)
    mask |= NSWindowStyleMaskMiniaturizable;
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
    [NSApp activateIgnoringOtherApps:YES];
    [window makeKeyAndOrderFront:NSApp];
  });
  return (uintptr_t)window;
}

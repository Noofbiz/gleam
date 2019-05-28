// +build darwin
// +build 386 amd64
// +build !ios

#import <Cocoa/Cocoa.h>

#include <float.h>
#include <string.h>

@interface AppDelegate : NSObject<NSApplicationDelegate>
{
}
@end

@implementation AppDelegate
- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
}

- (void)applicationWillHide:(NSNotification *)aNotification {
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

uintptr_t newWindow(int width, int height, char* title) {
  __block NSWindow* window = NULL;
  dispatch_sync(dispatch_get_main_queue(), ^{
    NSString* name = [[NSString alloc] initWithUTF8String:title];
    NSRect frame = NSMakeRect(0, 0, (CGFloat)width, (CGFloat)height);
    NSWindow* window  = [[[NSWindow alloc] initWithContentRect:frame
                      styleMask:NSWindowStyleMaskTitled|NSWindowStyleMaskClosable
                      backing:NSBackingStoreBuffered
                      defer:NO] autorelease];
    [window setBackgroundColor:[NSColor blueColor]];
    [window makeKeyAndOrderFront:NSApp];
    window.title = name;
  });
  return (uintptr_t)window;
}

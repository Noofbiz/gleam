// +build darwin
// +build 386 amd64
// +build !ios

#import <Cocoa/Cocoa.h>

#include <float.h>
#include <string.h>

int initializeGleam();
int appDone();
int setTitle(uintptr_t w, char* title);
int setFullScreen(uintptr_t w);
int resize(uintptr_t w, int width, int height, int x, int y, bool resizable);
int closeWindow(uintptr_t w);
int setBgColor(uintptr_t w, uint32_t red, uint32_t green, uint32_t blue, uint32_t alpha);
uintptr_t newWindow(int width, int height, int x, int y, char* title,
  bool resizable, bool fullscreen, uint32_t red, uint32_t green, uint32_t blue, uint32_t alpha);

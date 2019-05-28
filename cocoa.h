// +build darwin
// +build 386 amd64
// +build !ios

#import <Cocoa/Cocoa.h>

#include <float.h>
#include <string.h>

int initializeGleam();
int appDone();
uintptr_t newWindow(int width, int height, char* title);

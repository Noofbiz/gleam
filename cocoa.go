// +build darwin
// +build 386 amd64
// +build !ios

package gleam

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework OpenGL

#include "cocoa.h"
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func initImpl() error {
	err := C.initializeGleam()
	if err != 0 {
		return errors.New("unable to initalize gleam")
	}
	return nil
}

func appDoneImpl() error {
	err := C.appDone()
	if err != 0 {
		return errors.New("failed to complete app done")
	}
	return nil
}

func newWindowImpl(w *Window) error {
	title := C.CString(w.opts.Title)
	defer C.free(unsafe.Pointer(title))
	data := C.newWindow(C.int(w.opts.Width), C.int(w.opts.Height), title)
	w.data = uintptr(data)
	return nil
}

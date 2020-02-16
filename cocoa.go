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
	"image/color"
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func initImpl() error {
	err := C.initializeGleam()
	if err != 0 {
		return errors.New("COCOA: unable to initalize gleam")
	}
	return nil
}

func appDoneImpl() error {
	err := C.appDone()
	if err != 0 {
		return errors.New("COCOA: failed to complete app done")
	}
	return nil
}

func newWindowImpl(w *Window) error {
	title := C.CString(w.opts.Title)
	defer C.free(unsafe.Pointer(title))
	r, g, b, a := w.opts.BgColor.RGBA()
	data := C.newWindow(C.int(w.opts.Width), C.int(w.opts.Height),
		C.int(w.opts.X), C.int(w.opts.Y), title, C.bool(w.opts.Resizable),
		C.bool(w.opts.FullScreen), C.uint32_t(r), C.uint32_t(g), C.uint32_t(b), C.uint32_t(a))
	w.data = uintptr(data)
	return nil
}

func setTitleImpl(w *Window, s string) error {
	title := C.CString(s)
	defer C.free(unsafe.Pointer(title))
	err := C.setTitle(C.uintptr_t(w.data), title)
	if err != 0 {
		return errors.New("COCOA: unable to set title")
	}
	return nil
}

func setFullScreenImpl(w *Window) error {
	err := C.setFullScreen(C.uintptr_t(w.data))
	if err != 0 {
		return errors.New("COCOA: unable to set full screen")
	}
	return nil
}

func resizeImpl(w *Window, width, height int) error {
	err := C.resize(C.uintptr_t(w.data), C.int(width), C.int(height), C.int(w.opts.X), C.int(w.opts.Y), C.bool(w.opts.Resizable))
	if err != 0 {
		return errors.New("COCOA: unable to resize window")
	}
	return nil
}

func moveImpl(w *Window, x, y int) error {
	err := C.resize(C.uintptr_t(w.data), C.int(w.opts.Width), C.int(w.opts.Height), C.int(x), C.int(y), C.bool(w.opts.Resizable))
	if err != 0 {
		return errors.New("COCOA: unable to move window")
	}
	return nil
}

func setResizableImpl(w *Window, b bool) error {
	err := C.resize(C.uintptr_t(w.data), C.int(w.opts.Width), C.int(w.opts.Height), C.int(w.opts.X), C.int(w.opts.Y), C.bool(b))
	if err != 0 {
		if b {
			return errors.New("COCOA: unable to set window resizable")
		}
		return errors.New("COCOA: unable to unset window resizable")
	}
	return nil
}

func closeImpl(w *Window) error {
	err := C.closeWindow(C.uintptr_t(w.data))
	if err != 0 {
		return errors.New("COCOA: unable to close window")
	}
	return nil
}

func setBgColorImpl(w *Window, c color.Color) error {
	r, g, b, a := c.RGBA()
	err := C.setBgColor(C.uintptr_t(w.data), C.uint32_t(r), C.uint32_t(g), C.uint32_t(b), C.uint32_t(a))
	if err != 0 {
		return errors.New("COCOA: unable to set bg color of window")
	}
	return nil
}

//export gleamInitDoneSignal
func gleamInitDoneSignal() {
	initCh <- struct{}{}
}

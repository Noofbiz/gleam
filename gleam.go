package gleam

import (
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

// Main is the executable's entry point. Uses the smae app.App as gomobile to
// simplify design.
func Main(mainFunc func(a *App)) {
	//run mainfunc on a separate goroutine
	go func() {
		mainFunc(&App{})
		appDone()
	}()
	//initalize Gleam on main thread
	initGleam()
	//clean up Gleam when done
	cleanup()
}

// initGleam initalizes Gleam in an OS specific way. This is called during Main
// and blocks until the app quits.
func initGleam() error {
	if err := initImpl(); err != nil {
		return err
	}
	return nil
}

func appDone() error {
	if err := appDoneImpl(); err != nil {
		return err
	}
	return nil
}

func cleanup() error {
	return nil
}

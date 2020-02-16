package gleam

import (
	"log"
)

var (
	initCh chan struct{}
	errCh  chan error
)

// Run runs the app. It initalizes the window system, then runs the app's
// Preload, Main, and Cleanup functions, and then cleans up the OSes window system.
func Run(a App) {
	initCh = make(chan struct{})
	errCh = make(chan error)
	// error logging goroutine
	go func() {
		for {
			err := <-errCh
			if err != nil {
				switch currentErrorLevel {
				case ErrorLevelLog:
					log.Printf("gleam error: %v\n", err.Error())
				case ErrorLevelPanic:
					panic("gleam error: " + err.Error())
				case ErrorLevelIgnore:
					//do nothing
				}
			}
		}
	}()
	//run mainfunc on a separate goroutine
	go func() {
		<-initCh
		errCh <- a.Preload()
		errCh <- a.Main()
		errCh <- a.Cleanup()
		errCh <- appDone()
	}()
	//initalize Gleam on main thread
	errCh <- initGleam()
	//clean up Gleam when done
	errCh <- cleanup()
	close(errCh)
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

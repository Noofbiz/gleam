package gleam

// App is an application that can run gleam. The idea is that the application
// has five major pieces:
// (1) Before the application is initalized for the OS
// (2) After the App is initalized
// (3) A main run loop that handles drawing to the window surface
// (4) After the loop has ended but before the OS-specific runtime is finished
// (5) After the runtime is finished.
// Step (1) is everything done before calling gleam.Run(app), and step (5) is
// anything that comes after it returns. This interface provides functions for
// the rest of the steps. Preload is (2), Main is (3), and Cleanup is (4).
type App interface {
	Preload() error
	Main() error
	Cleanup() error
}

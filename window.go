package gleam

import "image/color"

// Window is a OS window.
type Window struct {
	data uintptr
	opts WindowOptions
}

// NewWindow creates a new window on the screen
func NewWindow(options WindowOptions) (*Window, error) {
	w := &Window{}
	w.opts = options
	if err := newWindowImpl(w); err != nil {
		return nil, err
	}
	return w, nil
}

// ChangeTitle sets the window's titlebar title to s
func (w *Window) ChangeTitle(s string) error {
	if err := setTitleImpl(w, s); err != nil {
		return err
	}
	w.opts.Title = s
	return nil
}

// Title returns the window's current title
func (w *Window) Title() string {
	return w.opts.Title
}

// Resize resizes the window to the passed width and height values.
// Resize makes full screen windows windowed, titled and bordered.
func (w *Window) Resize(width, height int) error {
	if err := resizeImpl(w, width, height); err != nil {
		return err
	}
	w.opts.Width = width
	w.opts.Height = height
	w.opts.FullScreen = false
	return nil
}

// Size returns the width and height of the window
func (w *Window) Size() (width int, height int) {
	width = w.opts.Width
	height = w.opts.Height
	return
}

// SetFullScreen sets the window to full screen if true and sets it to a bordered,
// titled window if false.
func (w *Window) SetFullScreen(b bool) error {
	if b {
		return setFullScreenImpl(w)
	}
	return w.Resize(w.Size())
}

// FullScreen returns a boolean indicating if the window is full screen
func (w *Window) FullScreen() bool {
	return w.opts.FullScreen
}

// Position returns the x and y coordinate, from the top left corner of the screen,
// of the window.
func (w *Window) Position() (x, y int) {
	x = w.opts.X
	y = w.opts.Y
	return
}

// Move moves the top left corner of the window to the x and y coordinates where
// 0,0 corresponds to the top left corner of the screen.
func (w *Window) Move(x, y int) error {
	if err := moveImpl(w, x, y); err != nil {
		return err
	}
	w.opts.X = x
	w.opts.Y = y
	return nil
}

// ChangeResizable changes the resizeable property of the window to the b
func (w *Window) ChangeResizable(b bool) error {
	if w.opts.Resizable == b {
		return nil //no need since they're the same
	}
	if err := setResizableImpl(w, b); err != nil {
		return err
	}
	w.opts.Resizable = b
	return nil
}

// Resizable returns a boolean indicating if the window is resizable
func (w *Window) Resizable() bool {
	return w.opts.Resizable
}

// Close closes the window and frees the associated memory. Do not make any more
// calls to the window after calling Close.
func (w *Window) Close() error {
	if err := closeImpl(w); err != nil {
		return err
	}
	return nil
}

// SetBgColor sets the background color of the window.
func (w *Window) SetBgColor(c color.Color) error {
	if err := setBgColorImpl(w, c); err != nil {
		return err
	}
	return nil
}

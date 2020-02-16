package gleam

import "image/color"

// WindowOptions are options used to create a window.
type WindowOptions struct {
	Height, Width int
	X, Y          int
	Title         string
	FullScreen    bool
	Resizable     bool
	BgColor       color.Color
}

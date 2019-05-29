package gleam

type WindowOptions struct {
	Height, Width                       int
	Title                               string
	FullScreen                          bool
	Resizable, Miniaturizable, Closable bool
	Titled, Bordered                    bool
}

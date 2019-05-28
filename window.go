package gleam

type Window struct {
	data uintptr
	opts WindowOptions
}

func (a *App) NewWindow(options WindowOptions) (*Window, error) {
	w := &Window{}
	w.opts = options
	if err := newWindowImpl(w); err != nil {
		return nil, err
	}
	return w, nil
}

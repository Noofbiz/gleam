package gleam

type App interface {
	Preload() error
	Main() error
	Cleanup() error
}

package gleam

type ErrorLevel uint

const (
	ErrorLevelLog ErrorLevel = iota
	ErrorLevelPanic
	ErrorLevelIgnore
)

var currentErrorLevel ErrorLevel

func SetErrorLevel(e ErrorLevel) {
	currentErrorLevel = e
}

package gleam

// ErrorLevel is the level of logging done when gleam encounters an error.
type ErrorLevel uint

const (
	//ErrorLevelLog prints the errors to the log
	ErrorLevelLog ErrorLevel = iota
	//ErrorLevelPanic causes the app to panic when it encounters an error
	ErrorLevelPanic
	//ErrorLevelIgnore ignores errors
	ErrorLevelIgnore
)

var currentErrorLevel ErrorLevel

// SetErrorLevel sets how errors are reported when gleam encounters one.
func SetErrorLevel(e ErrorLevel) {
	currentErrorLevel = e
}

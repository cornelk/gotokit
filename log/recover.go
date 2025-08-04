package log

import "fmt"

// RecoverPanics is a function that can be deferred to recover from panics.
// The panic will be logged as fatal message and the system exits with
// os.Exit(1), even if logging at FatalLevel is disabled.
func RecoverPanics() {
	if r := recover(); r != nil {
		logger := Must(New())
		if err, ok := r.(error); ok {
			logger.Fatal("Panic", Err(err))
		} else {
			logger.Fatal("Panic", String("info", fmt.Sprintf("%+v", r)))
		}
	}
}

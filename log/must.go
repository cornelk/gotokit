package log

// Must panics in case of an error.
func Must(logger *Logger, err error) *Logger {
	if err != nil {
		panic(err)
	}
	return logger
}

package error2

import (
	"fmt"
	"runtime"
	"strings"
)

func getCaller(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}

	return prettyCaller(file, line)
}

func prettyCaller(file string, line int) string {
	idx := strings.LastIndexByte(file, '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}

	idx = strings.LastIndexByte(file[:idx], '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}

	return fmt.Sprintf("%s:%d", file[idx+1:], line)
}

// Wrap returns an error that wraps err with given message.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("[%s] %s: %w", getCaller(2), message, err)
}

// Wrapf returns an error that wraps err with given format and args.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("[%s] %s: %w", getCaller(2), fmt.Sprintf(format, args...), err)
}

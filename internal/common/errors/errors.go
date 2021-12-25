package errors

import (
	"errors"

	"github.com/ansel1/merry/v2"
)

func init() {
	merry.SetStackCaptureEnabled(true)
}

func New(message string) error {
	return merry.New(message)
}

func Errorf(format string, args ...interface{}) error {
	return merry.Errorf(format, args...)
}

func StackTrace(err error) ([]string, bool) {
	if merry.HasStack(err) {
		return merry.FormattedStack(err), true
	}

	return []string{}, false
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

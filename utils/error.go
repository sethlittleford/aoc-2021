package utils

import (
	"errors"
)

// CheckErr is a helper function that checks if the given error
// is non-nil. If so, it appends custom error message(s) to the error and panics
func CheckErr(err error, message ...string) {
	if err != nil {
		s := err.Error()
		for _, msg := range message {
			s += " | " + msg
		}
		panic(errors.New(s))
	}
}

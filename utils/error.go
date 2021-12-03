package utils

import (
	"errors"
	"fmt"
)

// CheckErr is a helper function that checks if the given error
// is non-nil. If so, it appends the custom error message to the error and panics
func CheckErr(err error, message string) {
	if err != nil {
		panic(errors.New(fmt.Sprintf("%s | %s", err.Error(), message)))
	}
}

package errs

import (
	"fmt"
)

type MyCustomError struct {
	err error
}

func NewMyCustomError(err error) error {
	return &MyCustomError{
		err: err,
	}
}

func (e *MyCustomError) Error() string {
	name := getCurrentFunctionName()
	msg := fmt.Sprintf("%s (Func: %s)", e.err.Error(), name)
	return msg
}

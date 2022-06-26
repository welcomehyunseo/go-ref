package errs

import (
	"fmt"
)

type MyCustomError struct {
	err   error
	where string
}

func NewMyCustomError(err error) error {
	where := getCallerFunctionName()
	return &MyCustomError{
		err:   err,
		where: where,
	}
}

func (e *MyCustomError) Error() string {
	msg := fmt.Sprintf("%s (Where: %s)", e.err.Error(), e.Where())

	return msg
}

func (e *MyCustomError) Where() string {
	return e.where
}

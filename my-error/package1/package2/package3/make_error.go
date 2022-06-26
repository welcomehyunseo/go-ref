package package3

import (
	"errors"
	"myerror/errs"
)

func MakeError() error {
	err := errs.NewMyCustomError(errors.New("my custom error"))
	return err
}

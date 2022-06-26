package package2

import (
	"errors"
	"fmt"
	"myerror/errs"
)

func MakeError() {
	err := errs.NewMyCustomError(errors.New("my custom error"))
	fmt.Println(err.Error())
}

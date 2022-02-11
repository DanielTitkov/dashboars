package tasks

import (
	"fmt"
)

func newArgError(arg string) error {
	return fmt.Errorf("argument value for '%s' not found", arg)
}

func newArgTypeError(arg string) error {
	return fmt.Errorf("argument value for '%s' has wrong type", arg)
}

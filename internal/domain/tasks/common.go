package tasks

import (
	"fmt"
)

func newArgError(arg string) error {
	return fmt.Errorf("argument value for '%s' not found", arg)
}

func newArgTypeError(arg, exp, got string) error {
	return fmt.Errorf("argument value for '%s' has wrong type: expected %s but got %s", arg, exp, got)
}

func requireFloat(args map[string]interface{}, arg string) (float64, error) {
	v, ok := args[arg]
	if !ok {
		return 0, newArgError(arg)
	}

	ret, ok := v.(float64)
	if !ok {
		return 0, newArgTypeError(arg, "float64", fmt.Sprintf("%T", v))
	}

	return ret, nil
}

func maybeFloat(args map[string]interface{}, arg string) float64 {
	v, ok := args[arg]
	if !ok {
		return 0
	}

	ret, ok := v.(float64)
	if !ok {
		return 0
	}

	return ret
}

package errors

import (
	"errors"
	"fmt"
)

// Wrap an error into a wrapper. If nil passed, nil will be returned.
func Wrap(err error, wrapper error) error {
	if err != nil {
		if !errors.Is(err, wrapper) {
			return fmt.Errorf("%w: %v", wrapper, err)
		}
		return err
	}
	return nil
}

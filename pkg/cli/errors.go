package cli

import "fmt"

func wrap(err error, wrapper error) error {
	if err != nil {
		return fmt.Errorf("%w: %v", wrapper, err)
	}
	return nil
}

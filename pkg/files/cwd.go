package files

import (
	"os"

	"github.com/cardil/deviate/pkg/errors"
)

// ErrCannotChangeDirectory when cannot change directory.
var ErrCannotChangeDirectory = errors.New("cannot change directory")

// WithinDirectory executes given function within directory.
func WithinDirectory(path string, fn func() error) error {
	currentWD, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, ErrCannotChangeDirectory)
	}
	err = os.Chdir(path)
	if err != nil {
		return errors.Wrap(err, ErrCannotChangeDirectory)
	}
	defer func() {
		_ = os.Chdir(currentWD)
	}()
	return fn()
}

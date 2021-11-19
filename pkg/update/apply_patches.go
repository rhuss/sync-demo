package update

import (
	"os"
	"path"
	"strings"

	"github.com/cardil/deviate/pkg/errors"
	"github.com/cardil/deviate/pkg/log/color"
	"github.com/magefile/mage/sh"
)

func (o Operation) applyPatches() error {
	o.Println("- Apply patches if present")
	patchesDir := path.Join(o.Project.Path, "openshift", "patches")
	files, err := os.ReadDir(patchesDir)
	if err != nil {
		o.Println("-- No patches found")
		return nil //nolint:nilerr
	}
	o.Printf("-- Found %d patche(s)\n", len(files))
	for _, file := range files {
		if !file.Type().IsRegular() || !strings.HasSuffix(file.Name(), ".patch") {
			continue
		}
		filePath := path.Join(patchesDir, file.Name())
		o.Printf("-- Applying %s\n", color.Blue(filePath))

		// TODO: Consider rewriting this to Go native code instead shell invocation.
		err = withWorkingDirectory(o.Project.Path, func() error {
			return errors.Wrap(sh.RunV("git", "apply", filePath),
				ErrUpdateFailed)
		})
		if err != nil {
			return err
		}
	}

	return runSteps([]step{
		o.commitChanges(":fire: Apply carried patches"),
	})
}

func withWorkingDirectory(path string, fn func() error) error {
	currentWD, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, ErrUpdateFailed)
	}
	err = os.Chdir(path)
	if err != nil {
		return errors.Wrap(err, ErrUpdateFailed)
	}
	defer func() {
		_ = os.Chdir(currentWD)
	}()
	return fn()
}

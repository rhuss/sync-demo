package main

import (
	"os"

	"github.com/cardil/deviate/internal/cmd"
)

var (
	exitFunc = os.Exit    //nolint:gochecknoglobals
	opts     []cmd.Option //nolint:gochecknoglobals
)

func main() {
	exitFunc(cmd.Main(opts...))
}

package cmd

import (
	"os"
	"path"

	"github.com/cardil/deviate/pkg/cli"
	"github.com/cardil/deviate/pkg/metadata"
	"github.com/spf13/cobra"
)

func addFlags(root *cobra.Command, opts *cli.Options) {
	fl := root.PersistentFlags()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := path.Join(wd, ".deviate.yaml")
	fl.StringVar(&opts.ConfigPath, "config", config,
		metadata.Name+" configuration file")
}

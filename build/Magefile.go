//go:build mage
// +build mage

package main

import (
	"github.com/cardil/deviate/pkg/metadata"

	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/artifact"
	"github.com/wavesoftware/go-magetasks/pkg/artifact/platform"
	"github.com/wavesoftware/go-magetasks/pkg/checks"
	"github.com/wavesoftware/go-magetasks/pkg/git"
)

// Default target is set to binary.
//
//goland:noinspection GoUnusedGlobalVariable
var Default = magetasks.Build // nolint:deadcode,gochecknoglobals

func init() { //nolint:gochecknoinits
	bin := artifact.Binary{
		Metadata:  config.Metadata{Name: "deviate"},
		Platforms: []artifact.Platform{{OS: platform.Linux, Architecture: platform.AMD64}},
	}

	magetasks.Configure(config.Config{
		Version: &config.Version{
			Path:     metadata.VersionPath(),
			Resolver: git.NewVersionResolver(),
		},
		Artifacts: []config.Artifact{bin},
		Checks:    []config.Task{checks.GolangCiLint(withVersion("v1.57.2"))},
	})
}

func withVersion(v string) checks.GolangCiLintParam {
	return func(opts *checks.GolangCiLintOptions) {
		opts.Version = v
	}
}

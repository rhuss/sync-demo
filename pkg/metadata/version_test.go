package metadata_test

import (
	"testing"

	"github.com/cardil/deviate/pkg/metadata"
	"gotest.tools/v3/assert"
)

func TestVersion(t *testing.T) {
	assert.Check(t, metadata.Version != "v0.0.0")
}

func TestVersionPath(t *testing.T) {
	assert.Check(t, metadata.VersionPath() != "")
}

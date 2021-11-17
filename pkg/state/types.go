package state

import (
	"context"

	"github.com/cardil/deviate/pkg/config"
	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/log"
)

// State represents a state of running tool.
type State struct {
	*config.Config
	git.Repository
	context.Context
	log.Logger
	cancel context.CancelFunc
}

package state

import (
	"context"

	"github.com/cardil/deviate/pkg/config"
	"github.com/cardil/deviate/pkg/log"
	"github.com/go-git/go-git/v5"
)

// State represents a state of running tool.
type State struct {
	*git.Repository
	*config.Config
	context.Context
	log.Logger
	cancel context.CancelFunc
}

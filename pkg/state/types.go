package state

import (
	"context"

	"github.com/openshift-knative/deviate/pkg/config"
	"github.com/openshift-knative/deviate/pkg/config/git"
	"github.com/openshift-knative/deviate/pkg/log"
)

// State represents a state of running tool.
type State struct {
	*config.Config
	*config.Project
	git.Repository
	context.Context
	log.Logger
	cancel context.CancelFunc
}

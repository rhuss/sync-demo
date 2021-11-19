package git

import (
	"context"

	"github.com/cardil/deviate/pkg/config"
	gitv5 "github.com/go-git/go-git/v5"
)

// Repository is an implementation of git repository using Golang library.
type Repository struct {
	*gitv5.Repository
	config.Project
	context.Context
}

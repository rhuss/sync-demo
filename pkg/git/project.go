package git

import (
	"errors"
	"fmt"

	"github.com/cardil/deviate/pkg/config"
	"github.com/cardil/deviate/pkg/state"
	gitv5 "github.com/go-git/go-git/v5"
)

var (
	// ErrNotGitRepo when target isn't a git repository.
	ErrNotGitRepo = errors.New("not a git repository")

	// ErrRemoteOperationFailed when remote git repository operation failed.
	ErrRemoteOperationFailed = errors.New("remote operation failed")
)

// New creates a new Project from regular config.Project.
func New(project config.Project, state state.State) (Project, error) {
	r, err := gitv5.PlainOpen(project.Path)
	if err != nil {
		return Project{}, fmt.Errorf("%s - %w: %v", project.Path, ErrNotGitRepo, err)
	}
	return Project{
		Project: project,
		repo:    r,
		state:   state,
	}, nil
}

// Project is a project with Git information attached.
type Project struct {
	config.Project
	state state.State
	repo  *gitv5.Repository
}

// Repository returns a Git repository implementation.
func (p Project) Repository() *Repository {
	return &Repository{Repository: p.repo, Context: p.state.Context}
}

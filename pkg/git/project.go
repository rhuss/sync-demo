package git

import (
	"errors"
	"fmt"

	"github.com/cardil/deviate/pkg/config"
	"github.com/go-git/go-git/v5"
)

// ErrNotGitRepo when target isn't a git repository.
var ErrNotGitRepo = errors.New("not a git repository")

type Project struct {
	config.Project
	repo *git.Repository
}

func New(project config.Project) (Project, error) {
	r, err := git.PlainOpen(project.Path)
	if err != nil {
		return Project{}, fmt.Errorf("%s - %w: %v", project.Path, ErrNotGitRepo, err)
	}
	return Project{
		Project: project,
		repo:    r,
	}, nil
}

func (p Project) Repository() *git.Repository {
	return p.repo
}

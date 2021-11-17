package git

import (
	"context"
	"regexp"
	"strings"

	"github.com/cardil/deviate/pkg/errors"
	gitv5 "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/mitchellh/go-homedir"
	sshagent "github.com/xanzy/ssh-agent"
)

const (
	httpURL  = "http://"
	httpsURL = "https://"
)

// https://regex101.com/r/wsAYas/1
var gitAddressRe = regexp.MustCompile(`^(?:git://)?(?:([^@]+)@)?[^:]+:.+\.git$`)

// Repository is an implementation of git repository using Golang library.
type Repository struct {
	*gitv5.Repository
	context.Context
}

func (r Repository) ListRemote(url string) ([]*plumbing.Reference, error) {
	remote := gitv5.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{url},
	})

	opts := &gitv5.ListOptions{}
	if isSSH(url) {
		auth, err := sshAuth(url)
		if err != nil {
			return nil, err
		}
		opts.Auth = auth
	}
	refs, err := remote.ListContext(r.Context, opts)
	if err != nil {
		return nil, errors.Wrap(err, ErrRemoteOperationFailed)
	}
	return refs, nil
}

func sshAuth(url string) (ssh.AuthMethod, error) { //nolint:ireturn
	var auth ssh.AuthMethod
	var err error
	if sshagent.Available() {
		user := ""
		if matches := gitAddressRe.FindStringSubmatch(url); matches != nil {
			user = matches[1]
		}
		auth, err = ssh.NewSSHAgentAuth(user)
		if err != nil {
			return nil, errors.Wrap(err, ErrRemoteOperationFailed)
		}
		return auth, nil
	}
	var idRsa string
	idRsa, err = homedir.Expand("~/.ssh/id_rsa")
	if err != nil {
		return nil, errors.Wrap(err, ErrRemoteOperationFailed)
	}
	auth, err = ssh.NewPublicKeysFromFile("git", idRsa, "")
	if err != nil {
		return nil, errors.Wrap(err, ErrRemoteOperationFailed)
	}
	return auth, nil
}

func isSSH(url string) bool {
	return !(strings.HasPrefix(url, httpsURL) ||
		strings.HasPrefix(url, httpURL))
}

func (r Repository) Remote(name string) (string, error) {
	remote, err := r.Repository.Remote(name)
	if err != nil {
		return "", errors.Wrap(err, ErrRemoteOperationFailed)
	}
	return remote.Config().URLs[0], nil
}

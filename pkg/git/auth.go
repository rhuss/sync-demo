package git

import (
	"regexp"
	"strings"

	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/mitchellh/go-homedir"
	sshagent "github.com/xanzy/ssh-agent"
)

const (
	httpURL  = "http://"
	httpsURL = "https://"
)

// https://regex101.com/r/wsAYas/1
var gitAddressRe = regexp.MustCompile(`^(?:git://)?(?:([^@]+)@)?[^:]+:.+\.git$`)

func authentication(remote git.Remote) (ssh.AuthMethod, error) { //nolint:ireturn
	var auth ssh.AuthMethod
	if isHTTP(remote.URL) {
		return auth, nil
	}
	var err error
	if sshagent.Available() {
		user := ""
		if matches := gitAddressRe.FindStringSubmatch(remote.URL); matches != nil {
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

func isHTTP(url string) bool {
	return strings.HasPrefix(url, httpsURL) ||
		strings.HasPrefix(url, httpURL)
}

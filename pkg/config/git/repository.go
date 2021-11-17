package git

import "github.com/go-git/go-git/v5/plumbing"

// RemoteLister will list references of a GIT repository remote.
type RemoteLister interface {
	ListRemote(url string) ([]*plumbing.Reference, error)
}

// RemoteURLInformer will return a URL of a remote or error if such remote
// do not exist.
type RemoteURLInformer interface {
	Remote(name string) (string, error)
}

// Repository contains operations on underlying GIT repo.
type Repository interface {
	RemoteLister
	RemoteURLInformer
}

package update

import (
	"fmt"
	"regexp"

	"github.com/cardil/deviate/pkg/errors"
	"github.com/wavesoftware/go-magetasks/pkg/output/color"
)

func (o Operation) mirrorBranches() error {
	o.Println(">>> Check if there's an upstream release we need " +
		"to mirror downstream")

	missing, err := o.findMissingDownstreamReleases()
	if err != nil {
		return err
	}
	if len(missing) > 0 {
		o.Printf(">> Found missing releases: %s\n", color.Blue(fmt.Sprintf("%+q", missing)))
		for _, rel := range missing {
			err = o.mirrorRelease(rel)
			if err != nil {
				return err
			}
		}
	} else {
		o.Println(">> No missing releases found")
	}
	return nil
}

type release struct {
	major, minor string
}

func (r release) String() string {
	return r.major + "." + r.minor
}

func (o Operation) findMissingDownstreamReleases() ([]release, error) {
	var upstreamReleases, downstreamReleases []release
	var err error
	downstreamReleases, err = o.listReleases(false)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateFailed)
	}
	upstreamReleases, err = o.listReleases(true)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateFailed)
	}

	missing := make([]release, 0, len(upstreamReleases))
	for _, candidate := range upstreamReleases {
		found := false
		for _, downstreamRelease := range downstreamReleases {
			if candidate == downstreamRelease {
				found = true
				break
			}
		}
		if !found {
			missing = append(missing, candidate)
		}
	}

	return missing, nil
}

func (o Operation) listReleases(upstream bool) ([]release, error) {
	url := o.Config.Downstream
	re := regexp.MustCompile(o.Config.Branches.Searches.DownstreamReleases)
	if upstream {
		url = o.Config.Upstream
		re = regexp.MustCompile(o.Config.Branches.Searches.UpstreamReleases)
	}

	refs, err := o.Repository.ListRemote(url)
	if err != nil {
		return nil, errors.Wrap(err, ErrUpdateFailed)
	}

	releases := make([]release, 0)

	for _, ref := range refs {
		name := ref.Name()
		if name.IsBranch() {
			branch := name.Short()
			if matches := re.FindStringSubmatch(branch); matches != nil {
				version := release{matches[1], matches[2]}
				releases = append(releases, version)
			}
		}
	}

	return releases, nil
}

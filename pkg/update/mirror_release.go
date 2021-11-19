package update

import (
	"github.com/cardil/deviate/pkg/config/git"
	"github.com/cardil/deviate/pkg/errors"
	"github.com/cardil/deviate/pkg/log/color"
	"github.com/cardil/deviate/pkg/state"
)

func (o Operation) mirrorRelease(rel release) error {
	return runSteps([]step{
		o.createNewRelease(rel),
		o.addForkFiles,
		o.applyPatches,
		o.switchToMain,
		o.pushRelease(rel),
	})
}

func (o Operation) createNewRelease(rel release) step {
	o.Printf("- Creating new release: %s\n", color.Blue(rel.String()))
	upstream := git.Remote{Name: "upstream", URL: o.Config.Upstream}
	cnr := createNewRelease{State: o.State, rel: rel, remote: upstream}
	return cnr.step
}

func (o Operation) pushRelease(rel release) step {
	return func() error {
		o.Printf("- Publishing release: %s\n", color.Blue(rel.String()))
		branch, err := rel.Name(o.Config.ReleaseTemplates.Downstream)
		if err != nil {
			return err
		}
		pr := pushBranch{State: o.State, branch: branch}
		return runSteps(pr.steps())
	}
}

type createNewRelease struct {
	state.State
	rel    release
	remote git.Remote
}

func (r createNewRelease) step() error {
	upstreamBranch, err := r.rel.Name(r.Config.ReleaseTemplates.Upstream)
	if err != nil {
		return err
	}
	downstreamBranch, err := r.rel.Name(r.Config.ReleaseTemplates.Downstream)
	if err != nil {
		return err
	}
	return runSteps([]step{
		r.fetch,
		r.checkoutAsNewRelease(upstreamBranch, downstreamBranch),
	})
}

func (r createNewRelease) fetch() error {
	return errors.Wrap(r.Repository.Fetch(r.remote), ErrUpdateFailed)
}

func (r createNewRelease) checkoutAsNewRelease(upstreamBranch, downstreamBranch string) step {
	return func() error {
		return errors.Wrap(
			r.Repository.Checkout(r.remote, upstreamBranch).As(downstreamBranch),
			ErrUpdateFailed)
	}
}

type pushBranch struct {
	state.State
	branch string
}

func (p pushBranch) steps() []step {
	return []step{
		p.push,
		p.delete,
	}
}

func (p pushBranch) push() error {
	if p.DryRun {
		p.Logger.Println(color.Yellow("- Skipping push, because of dry run"))
		return nil
	}
	return errors.Wrap(p.Repository.PushRelease(p.branch), ErrUpdateFailed)
}

func (p pushBranch) delete() error {
	return errors.Wrap(p.Repository.DeleteBranch(p.branch), ErrUpdateFailed)
}

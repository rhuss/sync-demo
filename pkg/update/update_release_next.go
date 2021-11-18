package update

func (o Operation) updateReleaseNext() error {
	return runSteps([]step{
		o.resetReleaseNext,
		o.removeGithubWorkflows,
		o.addForkFiles,
		o.applyPatches,
	})
}

package update

func (o Operation) updateReleaseNext() error {
	steps := []step{
		o.resetReleaseNext,
		o.removeGithubWorkflows,
		o.addForkFiles,
		o.applyPatches,
	}

	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	return nil
}

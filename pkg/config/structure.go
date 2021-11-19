package config

// Config for a deviate to operate.
type Config struct {
	Upstream                   string `valid:"required"`
	Downstream                 string `valid:"required"`
	SynchronizeTags            bool
	DryRun                     bool
	GithubWorkflowsRemovalGlob string `valid:"required"`
	Branches
}

// Branches holds configuration for branches.
type Branches struct {
	Main        string `valid:"required"`
	ReleaseNext string `valid:"required"`
	SynchCI     string `valid:"required"`
	ReleaseTemplates
	Searches
}

// ReleaseTemplates contains templates for release names.
type ReleaseTemplates struct {
	Upstream   string `valid:"required"`
	Downstream string `valid:"required"`
}

// Searches contains regular expressions used to search for branches.
type Searches struct {
	UpstreamReleases   string `valid:"required"`
	DownstreamReleases string `valid:"required"`
}

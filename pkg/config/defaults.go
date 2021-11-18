package config

// newDefaults creates a new default configuration.
func newDefaults() Config {
	const (
		releaseTemplate = "release-{{ .Major }}.{{ .Minor }}"
		releaseSearch   = `release-(\d+)\.(\d+)`
	)
	return Config{
		Branches: Branches{
			Main:        "main",
			ReleaseNext: "release-next",
			ReleaseTemplates: ReleaseTemplates{
				Upstream:   releaseTemplate,
				Downstream: releaseTemplate,
			},
			Searches: Searches{
				UpstreamReleases:   releaseSearch,
				DownstreamReleases: releaseSearch,
			},
		},
	}
}

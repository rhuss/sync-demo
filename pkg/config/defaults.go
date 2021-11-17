package config

// newDefaults creates a new default configuration.
func newDefaults() Config {
	return Config{
		Branches: Branches{
			Main:            "main",
			ReleaseNext:     "release-next",
			ReleaseTemplate: "release-{{ .Version.Major }}-{{ .Version.Minor }}",
			Searches: Searches{
				UpstreamReleases:   `release-(\d+)\.(\d+)`,
				DownstreamReleases: `release-(\d+)\.(\d+)`,
			},
		},
	}
}

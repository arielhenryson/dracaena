package config

// DockerConfig is the configuration for docker runner
type DockerConfig struct {
	// the root folder inside the container
	// where the job will run
	Workspace string
}

// MainDockerConfig return pointer to DockerConfig
func MainDockerConfig() *DockerConfig {
	conf := DockerConfig{
		Workspace: "/app",
	}

	return &conf
}

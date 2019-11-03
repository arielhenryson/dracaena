package createCommand

import "github.com/arielhenryson/dracaena/internal/runners/docker/config"

// CreateCommand make user command able to run on sh
func CreateCommand(commands []string) []string  {
	dockerConfig := config.MainDockerConfig()

	command := "cd " + dockerConfig.Workspace

	for _, element := range commands {
		command += "&&" + element
	}

	return []string{"sh", "-c", command}
}
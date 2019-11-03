package createContainer

import (
	"context"
	"github.com/arielhenryson/dracaena/internal/runners/docker/config"
	"github.com/arielhenryson/dracaena/internal/runners/docker/helpers/createCommand"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// CreateContainer create docker container and return the id
func CreateContainer(cli *client.Client, ctx context.Context,
	image string,
	volumeName string,
	commands []string,
	containerName string,
	networkID string,
) (string, error)  {
	dockerConfig := config.MainDockerConfig()

	// format the commands to run on sh
	command := createCommand.CreateCommand(commands)

	// this wil connect to this container the same volume we are using
	// for all jobs so they can share data and also will use for cache
	// between workflow execution
	vol := mount.Mount{
		Type:   mount.TypeVolume,
		Source: volumeName,
		Target: dockerConfig.Workspace,
	}

	// this will connect the container to workflow network so container
	// will be able to talk with each other by containerName.networkName
	workflowNetwork := map[string]*network.EndpointSettings{
		networkID: &network.EndpointSettings{
			NetworkID: networkID,
		},
	}

	jobContainer, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: image,
			// this command prevent the container from exit
			// Cmd: strings.Fields("tail -f /dev/null"),
			Cmd: command,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{ vol },
		},
		&network.NetworkingConfig{
			EndpointsConfig: workflowNetwork,
		},
		containerName,
	)


	if err != nil {
		panic(err)
	}

	err = cli.ContainerStart(ctx, jobContainer.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	// send logs from container to the main log system
	go printLogs(cli, ctx, jobContainer.ID)

	// prevent the app from progress until container finish is work
	_, _ = waitContainerToFinish(cli, jobContainer.ID)

	// remove container after its finish
	_ = removeDoneContainer(ctx, cli, jobContainer.ID)

	return jobContainer.ID, nil
}

func removeDoneContainer(ctx context.Context, cli *client.Client, containerID string) error  {
	err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks: false,
		Force: false,
	})


	return err
}

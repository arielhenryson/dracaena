package createContainer

import (
	"context"
	"github.com/docker/docker/client"
)

// waitContainerToFinish prevent the app from progress until container finish is work
func waitContainerToFinish(cli *client.Client, containerID string) (int64, error) {
	// for now, we just wait for the app to exit and prints logs afterwards
	return cli.ContainerWait(context.Background(), containerID)
}
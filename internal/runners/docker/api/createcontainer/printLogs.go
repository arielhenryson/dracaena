package createContainer

import (
	"bufio"
	"context"
	"github.com/arielhenryson/dracaena/pkg/logger"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// printLogs print the logs of the container
func printLogs(cli *client.Client, ctx context.Context, containerID string) {
	reader, err := cli.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
	})
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		logger.Log(scanner.Text())
	}
}
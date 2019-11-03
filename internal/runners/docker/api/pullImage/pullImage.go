package pullImage

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

// PullImage pull image from docker-hub
func PullImage(ctx context.Context, cli *client.Client  ,image string)  {
	reader, err := cli.ImagePull(ctx, "docker.io/library/" + image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, reader)
}

package createVolume

import (
	"context"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// CreateVolume create volume to use for workflow
func CreateVolume(cli *client.Client, ctx context.Context, name string) (string, error)  {
	config := volume.VolumesCreateBody{
		Driver:     "local",
		DriverOpts: nil,
		Labels:     nil,
		Name:       name,
	}

	vol, _ := cli.VolumeCreate(ctx, config)

	return vol.Name, nil
}